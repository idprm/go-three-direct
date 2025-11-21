package entity

import "time"

type Transaction struct {
	ID             uint64 `json:"id"`
	TransactionID  string `json:"transaction_id"`
	ServiceID      int    `json:"-"`
	Service        Service
	Msisdn         string    `json:"msisdn"`
	SubmitedID     string    `json:"submited_id"`
	Keyword        string    `json:"keyword"`
	Adnet          string    `json:"adnet"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	StatusCode     int       `json:"status_code"`
	StatusDetail   string    `json:"status_detail"`
	Subject        string    `json:"subject"`
	DrStatus       string    `json:"dr_status"`
	DrStatusDetail string    `json:"dr_status_detail"`
	IpAddress      string    `json:"ip_address"`
	Payload        string    `json:"payload"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (e *Transaction) GetId() uint64 {
	return e.ID
}

func (e *Transaction) GetTransactionId() string {
	return e.TransactionID
}

func (e *Transaction) GetServiceId() int {
	return e.ServiceID
}

func (e *Transaction) GetMsisdn() string {
	return e.Msisdn
}

func (e *Transaction) GetSubmitedId() string {
	return e.SubmitedID
}

func (e *Transaction) GetKeyword() string {
	return e.Keyword
}

func (e *Transaction) GetAdnet() string {
	return e.Adnet
}

func (e *Transaction) GetAmount() float64 {
	return e.Amount
}

func (e *Transaction) GetStatus() string {
	return e.Status
}

func (e *Transaction) GetStatusCode() int {
	return e.StatusCode
}

func (e *Transaction) GetStatusDetail() string {
	return e.StatusDetail
}

func (e *Transaction) GetSubject() string {
	return e.Subject
}

func (e *Transaction) GetDrStatus() string {
	return e.DrStatus
}

func (e *Transaction) GetDrStatusDetail() string {
	return e.DrStatusDetail
}

func (e *Transaction) GetIpAddress() string {
	return e.IpAddress
}

func (e *Transaction) GetPayload() string {
	return e.Payload
}

func (e *Transaction) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *Transaction) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}
