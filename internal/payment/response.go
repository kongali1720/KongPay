package payment

type TransferResponse struct {
	Success     bool   `json:"success"`
	ReferenceNo string `json:"reference_no"`
	Status      string `json:"status"`
	Message     string `json:"message,omitempty"`
}
