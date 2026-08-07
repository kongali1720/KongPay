package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kongali1720/KongPay/internal/payment/security"
)

// Webhook menangani callback dari payment provider (Midtrans/Xendit/QRIS).
// WAJIB verifikasi signature/token SEBELUM memproses status apapun -
// tanpa ini, siapapun bisa memalsukan notifikasi pembayaran sukses.
func (h *PaymentHandler) Webhook(c *gin.Context) {
	providerType := c.Query("provider")
	if providerType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provider parameter required"})
		return
	}

	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Printf("❌ Invalid webhook payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// --- VERIFIKASI WAJIB, per provider ---
	switch providerType {
	case "midtrans":
		orderID, _ := payload["order_id"].(string)
		statusCode, _ := payload["status_code"].(string)
		grossAmount, _ := payload["gross_amount"].(string)
		signatureKey, _ := payload["signature_key"].(string)

		if err := security.VerifyMidtransSignature(orderID, statusCode, grossAmount, signatureKey); err != nil {
			log.Printf("🚨 WEBHOOK REJECTED (midtrans): %v - order_id=%s", err, orderID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "signature verification failed"})
			return
		}

	case "xendit":
		callbackToken := c.GetHeader("X-CALLBACK-TOKEN")
		if err := security.VerifyXenditToken(callbackToken); err != nil {
			log.Printf("🚨 WEBHOOK REJECTED (xendit): %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid callback token"})
			return
		}

	default:
		log.Printf("🚨 WEBHOOK REJECTED: unknown/unsupported provider=%s", providerType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported provider"})
		return
	}
	// --- Verifikasi lolos, baru boleh proses ---

	transactionID, _ := payload["transaction_id"].(string)
	if transactionID == "" {
		// Midtrans pakai field "order_id", bukan "transaction_id"
		transactionID, _ = payload["order_id"].(string)
	}
	status, _ := payload["status"].(string)

	log.Printf("📨 Webhook verified: TX=%s, Status=%s, Provider=%s", transactionID, status, providerType)

	if err := h.txService.UpdateTransactionStatus(c.Request.Context(), transactionID, "SUCCESS"); err != nil {
		log.Printf("❌ Failed to update status: %v", err)
	}

	tx, err := h.txService.GetTransaction(c.Request.Context(), transactionID)
	if err != nil {
		log.Printf("❌ Error getting transaction: %v", err)
		c.JSON(http.StatusOK, gin.H{"status": "webhook_processed"})
		return
	}
	if tx == nil {
		log.Printf("❌ Transaction NOT found in DB: %s", transactionID)
		c.JSON(http.StatusOK, gin.H{"status": "webhook_processed"})
		return
	}

	log.Printf("💰 Settlement: %s Amount: %.2f %s", transactionID, tx.Amount, tx.Currency)
	if err := h.txService.TriggerSettlementDirect(c.Request.Context(), transactionID, tx.Amount, tx.Currency, tx.CustomerID, tx.MerchantID); err != nil {
		log.Printf("❌ Settlement failed: %v", err)
	} else {
		log.Printf("✅ Settlement triggered: %s", transactionID)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "webhook_processed",
		"message": "Webhook processed successfully",
	})
}
