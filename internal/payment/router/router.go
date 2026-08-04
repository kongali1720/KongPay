package router

import (
    "context"
    "fmt"
    "sync"

    "github.com/kongali1720/KongPay/internal/payment/provider"
)

type PaymentRouter struct {
    providers map[provider.ProviderType]provider.PaymentProvider
    mu        sync.RWMutex
}

func NewPaymentRouter() *PaymentRouter {
    return &PaymentRouter{
        providers: make(map[provider.ProviderType]provider.PaymentProvider),
    }
}

func (r *PaymentRouter) Register(p provider.PaymentProvider) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.providers[p.Type()] = p
}

func (r *PaymentRouter) Route(ctx context.Context, req *provider.PaymentRequest) (*provider.PaymentResponse, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    providerType := r.selectProvider(req)

    p, exists := r.providers[providerType]
    if !exists {
        return nil, fmt.Errorf("provider %s not registered", providerType)
    }

    if !p.IsAvailable(ctx) {
        return r.fallback(ctx, req)
    }

    return p.Process(ctx, req)
}

func (r *PaymentRouter) selectProvider(req *provider.PaymentRequest) provider.ProviderType {
    switch req.Method {
    case "BANK_TRANSFER", "bank", "transfer":
        return provider.BankTransfer
    case "QRIS", "qris":
        return provider.QRIS
    case "CRYPTO", "crypto", "USDT", "USDC":
        return provider.Crypto
    default:
        return provider.BankTransfer
    }
}

func (r *PaymentRouter) fallback(ctx context.Context, req *provider.PaymentRequest) (*provider.PaymentResponse, error) {
    fallbacks := []provider.ProviderType{
        provider.BankTransfer,
        provider.QRIS,
        provider.Crypto,
    }

    for _, pt := range fallbacks {
        if p, exists := r.providers[pt]; exists && p.IsAvailable(ctx) {
            return p.Process(ctx, req)
        }
    }

    return nil, fmt.Errorf("no provider available")
}

func (r *PaymentRouter) GetProvider(pt provider.ProviderType) (provider.PaymentProvider, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    p, exists := r.providers[pt]
    return p, exists
}
