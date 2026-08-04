package listener

import (
    "context"
    "log"
)

type EthereumListener struct {
    rpcURL      string
    walletAddrs []string
    ctx         context.Context
    cancel      context.CancelFunc
}

func NewEthereumListener(rpcURL string) *EthereumListener {
    ctx, cancel := context.WithCancel(context.Background())
    return &EthereumListener{
        rpcURL: rpcURL,
        ctx:    ctx,
        cancel: cancel,
    }
}

func (e *EthereumListener) AddWallet(address string) {
    e.walletAddrs = append(e.walletAddrs, address)
}

func (e *EthereumListener) Start() error {
    log.Printf("Starting Ethereum listener for %d wallets", len(e.walletAddrs))
    go e.listen()
    return nil
}

func (e *EthereumListener) Stop() {
    e.cancel()
    log.Println("Ethereum listener stopped")
}

func (e *EthereumListener) listen() {
    <-e.ctx.Done()
}

func (e *EthereumListener) HandleTransaction(txHash string, from string, to string, amount string) error {
    log.Printf("Transaction detected: %s from %s to %s amount %s", txHash, from, to, amount)
    return nil
}
