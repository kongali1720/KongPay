package security

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidSignature = errors.New("invalid webhook signature")
	ErrMissingSignature = errors.New("webhook signature missing")
)

// VerifyMidtransSignature mencocokkan signature_key yang dikirim Midtrans.
// Formula resmi Midtrans: SHA512(order_id + status_code + gross_amount + ServerKey)
func VerifyMidtransSignature(orderID, statusCode, grossAmount, signatureKey string) error {
	if signatureKey == "" {
		return ErrMissingSignature
	}

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		return fmt.Errorf("MIDTRANS_SERVER_KEY tidak di-set")
	}

	raw := orderID + statusCode + grossAmount + serverKey
	hash := sha512.Sum512([]byte(raw))
	expected := hex.EncodeToString(hash[:])

	if !hmac.Equal([]byte(expected), []byte(signatureKey)) {
		return ErrInvalidSignature
	}
	return nil
}

// VerifyXenditToken mencocokkan header X-CALLBACK-TOKEN dengan token
// yang di-set di dashboard Xendit (Settings > Callbacks).
func VerifyXenditToken(receivedToken string) error {
	if receivedToken == "" {
		return ErrMissingSignature
	}

	expectedToken := os.Getenv("XENDIT_CALLBACK_TOKEN")
	if expectedToken == "" {
		return fmt.Errorf("XENDIT_CALLBACK_TOKEN tidak di-set")
	}

	if !hmac.Equal([]byte(receivedToken), []byte(expectedToken)) {
		return ErrInvalidSignature
	}
	return nil
}
