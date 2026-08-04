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
    db            *pgxpool.Pool
    mu            sync.RWMutex
    transactions  map[string]*SettlementTask
    workerCount   int
    taskQueue     chan *SettlementTask
}

func NewSettlementService(db *pgxpool.Pool) *SettlementService {
    if db == nil {
        log.Println("⚠️ SettlementService: DB is nil, settlements will NOT be saved to database!")
    } else {
        log.Println("✅ SettlementService: DB connected")
    }

    s := &SettlementService{
        db:            db,
        transactions: make(map[string]*SettlementTask),
        workerCount:   3,
        taskQueue:     make(chan *SettlementTask, 1000),
    }

    for i := 0; i < s.workerCount; i++ {
        go s.worker(i)
    }

    return s
}

func (s *SettlementService) AddTask(ctx context.Context, transactionID string, amount float64, currency string, customerID string, merchantID string) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    task := &SettlementTask{
        TransactionID: transactionID,
        Amount:        amount,
        Currency:      currency,
        CustomerID:    customerID,
        MerchantID:    merchantID,
        Status:        "PENDING",
        CreatedAt:     time.Now(),
    }

    s.transactions[transactionID] = task
    s.taskQueue <- task

    // Save to database (synchronous agar terlihat errornya)
    if s.db != nil {
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
            return err
        }
        log.Printf("✅ Settlement saved to DB: %s", transactionID)
    } else {
        log.Printf("⚠️ DB is nil, settlement NOT saved to database: %s", transactionID)
    }

    log.Printf("📥 Settlement task added: %s (Amount: %.2f %s)", transactionID, amount, currency)
    return nil
}

func (s *SettlementService) worker(id int) {
    log.Printf("🔄 Settlement worker %d started", id)

    for task := range s.taskQueue {
        log.Printf("🔄 Worker %d processing settlement: %s", id, task.TransactionID)

        if err := s.processSettlement(task); err != nil {
            log.Printf("❌ Settlement failed for %s: %v", task.TransactionID, err)
            task.Status = "FAILED"
            task.Error = err.Error()
        } else {
            log.Printf("✅ Settlement completed for %s", task.TransactionID)
            task.Status = "COMPLETED"
            task.CompletedAt = time.Now()
        }

        // Update DB status
        if s.db != nil {
            query := `
                UPDATE settlements 
                SET status = $1, completed_at = $2, error = $3
                WHERE transaction_id = $4
            `
            _, err := s.db.Exec(context.Background(), query,
                task.Status, task.CompletedAt, task.Error, task.TransactionID,
            )
            if err != nil {
                log.Printf("❌ Failed to update settlement status: %v", err)
            } else {
                log.Printf("✅ Settlement status updated in DB: %s -> %s", task.TransactionID, task.Status)
            }
        }

        s.mu.Lock()
        s.transactions[task.TransactionID] = task
        s.mu.Unlock()
    }
}

func (s *SettlementService) processSettlement(task *SettlementTask) error {
    // TODO: Implement real settlement logic
    time.Sleep(2 * time.Second)
    return nil
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
