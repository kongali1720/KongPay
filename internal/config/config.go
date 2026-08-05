package config

import (
    "os"
    "strconv"
    "strings"
)

type Config struct {
    // Server
    Port     string
    GinMode  string
    LogLevel string

    // Database
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string

    // Fiat Providers
    FiatConfig FiatConfig

    // Crypto
    CryptoConfig CryptoConfig
}

type FiatConfig struct {
    MidtransEnabled   bool
    MidtransServerKey string
    MidtransClientKey string
    MidtransBaseURL   string

    QRISEnabled   bool
    QRISMerchantID string
    QRISAPIKey     string
    QRISBaseURL    string
}

type CryptoConfig struct {
    ETHEnabled      bool
    ETHRPCURL       string
    ETHWSURL        string
    ETHChainID      int
    ETHMinConfirmations int

    BSCEnabled      bool
    BSCRPCURL       string
    BSCWSURL        string
    BSCChainID      int
    BSCMinConfirmations int

    PolygonEnabled  bool
    PolygonRPCURL   string
    PolygonWSURL    string
    PolygonChainID  int
    PolygonMinConfirmations int

    WalletHDSeed    string
    WalletPath      string
    WalletEncryptionKey string

    Tokens          map[string]string
}

func Load() *Config {
    return &Config{
        Port:     getEnv("PORT", "8080"),
        GinMode:  getEnv("GIN_MODE", "release"),
        LogLevel: getEnv("LOG_LEVEL", "info"),

        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBName:     getEnv("DB_NAME", "kongpay"),
        DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

        FiatConfig: FiatConfig{
            MidtransEnabled:   getBool("MIDTRANS_ENABLED", false),
            MidtransServerKey: getEnv("MIDTRANS_SERVER_KEY", ""),
            MidtransClientKey: getEnv("MIDTRANS_CLIENT_KEY", ""),
            MidtransBaseURL:   getEnv("MIDTRANS_BASE_URL", ""),
            QRISEnabled:       getBool("QRIS_ENABLED", false),
            QRISMerchantID:    getEnv("QRIS_MERCHANT_ID", ""),
            QRISAPIKey:        getEnv("QRIS_API_KEY", ""),
            QRISBaseURL:       getEnv("QRIS_BASE_URL", ""),
        },

        CryptoConfig: CryptoConfig{
            ETHEnabled:      getBool("ETH_ENABLED", false),
            ETHRPCURL:       getEnv("ETH_RPC_URL", ""),
            ETHWSURL:        getEnv("ETH_WS_URL", ""),
            ETHChainID:      getInt("ETH_CHAIN_ID", 1),
            ETHMinConfirmations: getInt("ETH_MIN_CONFIRMATIONS", 12),

            BSCEnabled:      getBool("BSC_ENABLED", false),
            BSCRPCURL:       getEnv("BSC_RPC_URL", ""),
            BSCWSURL:        getEnv("BSC_WS_URL", ""),
            BSCChainID:      getInt("BSC_CHAIN_ID", 56),
            BSCMinConfirmations: getInt("BSC_MIN_CONFIRMATIONS", 12),

            PolygonEnabled:  getBool("POLYGON_ENABLED", false),
            PolygonRPCURL:   getEnv("POLYGON_RPC_URL", ""),
            PolygonWSURL:    getEnv("POLYGON_WS_URL", ""),
            PolygonChainID:  getInt("POLYGON_CHAIN_ID", 137),
            PolygonMinConfirmations: getInt("POLYGON_MIN_CONFIRMATIONS", 64),

            WalletHDSeed:    getEnv("WALLET_HD_SEED", ""),
            WalletPath:      getEnv("WALLET_DERIVATION_PATH", "m/44'/60'/0'/0/0"),
            WalletEncryptionKey: getEnv("WALLET_ENCRYPTION_KEY", ""),
            Tokens: map[string]string{
                "USDT": getEnv("TOKEN_USDT_CONTRACT", ""),
                "USDC": getEnv("TOKEN_USDC_CONTRACT", ""),
                "DAI":  getEnv("TOKEN_DAI_CONTRACT", ""),
            },
        },
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}

func getBool(key string, fallback bool) bool {
    if value := os.Getenv(key); value != "" {
        return strings.ToLower(value) == "true"
    }
    return fallback
}

func getInt(key string, fallback int) int {
    if value := os.Getenv(key); value != "" {
        if i, err := strconv.Atoi(value); err == nil {
            return i
        }
    }
    return fallback
}
