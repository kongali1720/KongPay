package router

import (
    "context"
    "fmt"
    "sync"
    
    "kongpay/internal/payment/provider"
)

// PaymentRouter routes payments to appropriate providers
type PaymentRouter struct {
    providers map[provider.ProviderType]provider.PaymentProvider
    mu        sync.RWMutex
}

// NewPaymentRouter creates a new payment router
func NewPaymentRouter() *PaymentRouter {
    return &PaymentRouter{
        providers: make(map[provider.ProviderType]provider.PaymentProvider),
    }
}

// Register adds a payment provider
func (r *PaymentRouter) Register(p provider.PaymentProvider) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.providers[p.Type()] = p
}

// Route routes a payment request to appropriate provider
func (r *PaymentRouter) Route(ctx context.Context, req *provider.PaymentRequest) (*provider.PaymentResponse, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    // Select provider based on request
    providerType := r.selectProvider(req)
    
    p, exists := r.providers[providerType]
    if !exists {
        return nil, fmt.Errorf("provider %s not registered", providerType)
    }
    
    // Check provider availability
    if !p.IsAvailable(ctx) {
        // Try fallback
        return r.fallback(ctx, req)
    }
    
    return p.Process(ctx, req)
}

// selectProvider chooses the best provider for the request
func (r *PaymentRouter) selectProvider(req *provider.PaymentRequest) provider.ProviderType {
    switch req.Method {
    case "BANK_TRANSFER", "bank", "transfer":
        return provider.BankTransfer
    case "QRIS", "qris":
        return provider.QRIS
    case "CRYPTO", "crypto", "USDT", "USDC":
        return provider.Crypto
    case "VIRTUAL_ACCOUNT", "va":
        return provider.VirtualAccount
    default:
        return provider.BankTransfer
    }
}

// fallback tries other providers if primary fails
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

// GetProvider returns a provider by type
func (r *PaymentRouter) GetProvider(pt provider.ProviderType) (provider.PaymentProvider, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    p, exists := r.providers[pt]
    return p, exists
}
