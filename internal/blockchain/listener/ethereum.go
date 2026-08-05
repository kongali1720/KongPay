package listener

import (
    "context"
    "log"
    "math/big"
    "sync"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

type TransactionEvent struct {
    TxHash        string  `json:"tx_hash"`
    From          string  `json:"from"`
    To            string  `json:"to"`
    Amount        float64 `json:"amount"`
    Token         string  `json:"token"`
    BlockNumber   uint64  `json:"block_number"`
    Confirmations int     `json:"confirmations"`
    Status        string  `json:"status"`
    TransactionID string  `json:"transaction_id"`
}

type EthereumListener struct {
    client      *ethclient.Client
    wallets     map[string]string
    txChan      chan TransactionEvent
    ctx         context.Context
    cancel      context.CancelFunc
    mu          sync.RWMutex
    minConfirm  int
}

func NewEthereumListener(rpcURL string, minConfirm int) (*EthereumListener, error) {
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithCancel(context.Background())

    return &EthereumListener{
        client:     client,
        wallets:    make(map[string]string),
        txChan:     make(chan TransactionEvent, 100),
        ctx:        ctx,
        cancel:     cancel,
        minConfirm: minConfirm,
    }, nil
}

func (el *EthereumListener) AddWallet(address, transactionID string) {
    el.mu.Lock()
    defer el.mu.Unlock()
    el.wallets[address] = transactionID
    log.Printf("📝 Monitoring wallet: %s for TX: %s", address, transactionID)
}

func (el *EthereumListener) Start() error {
    log.Printf("🔍 Starting Ethereum listener with %d wallets", len(el.wallets))
    headers := make(chan *types.Header)
    sub, err := el.client.SubscribeNewHead(el.ctx, headers)
    if err != nil {
        return err
    }

    go func() {
        for {
            select {
            case <-el.ctx.Done():
                log.Println("🛑 Ethereum listener stopped")
                return
            case err := <-sub.Err():
                log.Printf("❌ Subscription error: %v", err)
                return
            case header := <-headers:
                el.processBlock(header)
            }
        }
    }()

    return nil
}

func (el *EthereumListener) processBlock(header *types.Header) {
    block, err := el.client.BlockByNumber(el.ctx, header.Number)
    if err != nil {
        log.Printf("❌ Failed to get block: %v", err)
        return
    }

    for _, tx := range block.Transactions() {
        // Get sender address
        from, err := el.getSender(tx)
        if err != nil {
            continue
        }

        to := tx.To()
        if to == nil {
            continue
        }

        toAddress := to.Hex()
        el.mu.RLock()
        transactionID, exists := el.wallets[toAddress]
        el.mu.RUnlock()

        if exists {
            amount := new(big.Float)
            amount.SetInt(tx.Value())
            amountFloat, _ := amount.Float64()
            amountFloat = amountFloat / 1e18

            log.Printf("💰 Transaction detected: %s -> %s amount: %f ETH", from.Hex(), toAddress, amountFloat)

            event := TransactionEvent{
                TxHash:        tx.Hash().Hex(),
                From:          from.Hex(),
                To:            toAddress,
                Amount:        amountFloat,
                Token:         "ETH",
                BlockNumber:   block.NumberU64(),
                Status:        "PENDING",
                TransactionID: transactionID,
            }

            el.txChan <- event
        }
    }
}

func (el *EthereumListener) getSender(tx *types.Transaction) (common.Address, error) {
    // Try to get sender from transaction
    signer := types.NewEIP155Signer(tx.ChainId())
    from, err := types.Sender(signer, tx)
    if err != nil {
        return common.Address{}, err
    }
    return from, nil
}

func (el *EthereumListener) GetTransactionChannel() chan TransactionEvent {
    return el.txChan
}

func (el *EthereumListener) Stop() {
    el.cancel()
    el.client.Close()
    log.Println("Ethereum listener stopped")
}

func (el *EthereumListener) CheckConfirmations(txHash string) (int, error) {
    tx, isPending, err := el.client.TransactionByHash(el.ctx, common.HexToHash(txHash))
    if err != nil {
        return 0, err
    }
    if isPending {
        return 0, nil
    }

    receipt, err := el.client.TransactionReceipt(el.ctx, tx.Hash())
    if err != nil {
        return 0, err
    }

    currentBlock, err := el.client.BlockNumber(el.ctx)
    if err != nil {
        return 0, err
    }

    confirmations := int(currentBlock - receipt.BlockNumber.Uint64())
    return confirmations, nil
}
