package services

import (
    "context"
    "log"
    "sync"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
)

type SettlementTask struct {
    TransactionID string    `json:"transaction_id"`
    Amount        float64   `json:"amount"`
    Currency      string    `json:"currency"`
    CustomerID    string    `json:"customer_id"`
    MerchantID    string    `json:"merchant_id"`
    Status        string    `json:"status"`
    CreatedAt     time.Time `json:"created_at"`
    CompletedAt   time.Time `json:"completed_at"`
    Error         string    `json:"error"`
}

type SettlementService struct {
    db           *pgxpool.Pool
    mu           sync.RWMutex
    transactions map[string]*SettlementTask
    workerCount  int
    taskQueue    chan *SettlementTask
}

func NewSettlementService(db *pgxpool.Pool) *SettlementService {
    s := &SettlementService{
        db:           db,
        transactions: make(map[string]*SettlementTask),
        workerCount:  3,
        taskQueue:    make(chan *SettlementTask, 100),
    }

    for i := 0; i < s.workerCount; i++ {
        go s.worker(i)
    }

    log.Println("✅ Settlement service initialized")
    return s
}

func (s *SettlementService) AddTask(ctx context.Context, transactionID string, amount float64, currency string, customerID string, merchantID string) error {
    log.Printf("📥 Adding settlement task: %s", transactionID)

    task := &SettlementTask{
        TransactionID: transactionID,
        Amount:        amount,
        Currency:      currency,
        CustomerID:    customerID,
        MerchantID:    merchantID,
        Status:        "PENDING",
        CreatedAt:     time.Now(),
    }

    s.mu.Lock()
    s.transactions[transactionID] = task
    s.mu.Unlock()

    s.taskQueue <- task
    log.Printf("✅ Task added to queue: %s", transactionID)

    // Save to database
    if s.db != nil {
        go s.saveToDB(ctx, task)
    }

    return nil
}

func (s *SettlementService) saveToDB(ctx context.Context, task *SettlementTask) {
    query := `
        INSERT INTO settlements (transaction_id, amount, currency, customer_id, merchant_id, status, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
    _, err := s.db.Exec(ctx, query,
        task.TransactionID, task.Amount, task.Currency,
        task.CustomerID, task.MerchantID, task.Status, task.CreatedAt,
    )
    if err != nil {
        log.Printf("❌ Failed to save settlement to DB: %v", err)
    } else {
        log.Printf("✅ Settlement saved to DB: %s", task.TransactionID)
    }
}

func (s *SettlementService) worker(id int) {
    log.Printf("🔄 Settlement worker %d started", id)

    for task := range s.taskQueue {
        log.Printf("🔄 Worker %d processing settlement: %s", id, task.TransactionID)

        // Simulate processing
        time.Sleep(2 * time.Second)

        task.Status = "COMPLETED"
        task.CompletedAt = time.Now()

        s.mu.Lock()
        s.transactions[task.TransactionID] = task
        s.mu.Unlock()

        // Update DB
        if s.db != nil {
            query := `
                UPDATE settlements 
                SET status = $1, completed_at = $2
                WHERE transaction_id = $3
            `
            _, err := s.db.Exec(context.Background(), query,
                task.Status, task.CompletedAt, task.TransactionID,
            )
            if err != nil {
                log.Printf("❌ Failed to update settlement status: %v", err)
            } else {
                log.Printf("✅ Settlement completed: %s", task.TransactionID)
            }
        }
    }
}

func (s *SettlementService) GetStatus(transactionID string) (*SettlementTask, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    if task, exists := s.transactions[transactionID]; exists {
        return task, nil
    }

    if s.db != nil {
        query := `
            SELECT transaction_id, amount, currency, customer_id, merchant_id, 
                   status, error, created_at, completed_at
            FROM settlements WHERE transaction_id = $1
        `
        var task SettlementTask
        var completedAt *time.Time
        err := s.db.QueryRow(context.Background(), query, transactionID).Scan(
            &task.TransactionID, &task.Amount, &task.Currency,
            &task.CustomerID, &task.MerchantID,
            &task.Status, &task.Error, &task.CreatedAt, &completedAt,
        )
        if err != nil {
            return nil, nil
        }
        if completedAt != nil {
            task.CompletedAt = *completedAt
        }
        return &task, nil
    }

    return nil, nil
}

func (s *SettlementService) GetStats() map[string]interface{} {
    s.mu.RLock()
    defer s.mu.RUnlock()

    total := len(s.transactions)
    pending := 0
    completed := 0
    failed := 0

    for _, task := range s.transactions {
        switch task.Status {
        case "PENDING":
            pending++
        case "COMPLETED":
            completed++
        case "FAILED":
            failed++
        }
    }

    return map[string]interface{}{
        "total":     total,
        "pending":   pending,
        "completed": completed,
        "failed":    failed,
    }
}
