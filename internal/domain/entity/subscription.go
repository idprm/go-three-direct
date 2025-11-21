package entity

import (
	"time"
)

type Subscription struct {
	ID            uint64 `json:"id"`
	ServiceID     int    `json:"service_id"`
	Service       Service
	Msisdn        string    `json:"msisdn"`
	Keyword       string    `json:"keyword"`
	Adnet         string    `json:"adnet"`
	LatestSubject string    `json:"latest_subject"`
	LatestStatus  string    `json:"latest_status"`
	Amount        float64   `json:"amount"`
	RenewalAt     time.Time `json:"renewal_at"`
	PurgeAt       time.Time `json:"purge_at"`
	UnsubAt       time.Time `json:"unsub_at"`
	ChargeAt      time.Time `json:"charge_at"`
	RetryAt       time.Time `json:"retry_at"`
	Success       uint      `json:"success"`
	IpAddress     string    `json:"ip_address"`
	IsRetry       bool      `json:"is_retry"`
	IsPurge       bool      `json:"is_purge"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (e *Subscription) GetId() uint64 {
	return e.ID
}

func (e *Subscription) GetServiceId() int {
	return e.ServiceID
}

func (e *Subscription) GetMsisdn() string {
	return e.Msisdn
}

func (e *Subscription) GetKeyword() string {
	return e.Keyword
}

func (e *Subscription) GetAdnet() string {
	return e.Adnet
}

func (e *Subscription) GetLatestSubject() string {
	return e.LatestSubject
}

func (e *Subscription) GetLatestStatus() string {
	return e.LatestStatus
}

func (e *Subscription) GetAmount() float64 {
	return e.Amount
}

func (e *Subscription) GetRenewalAt() time.Time {
	return e.RenewalAt
}

func (e *Subscription) GetPurgeAt() time.Time {
	return e.PurgeAt
}

func (e *Subscription) GetUnsubAt() time.Time {
	return e.UnsubAt
}

func (e *Subscription) GetChargeAt() time.Time {
	return e.ChargeAt
}

func (e *Subscription) GetRetryAt() time.Time {
	return e.RetryAt
}

func (e *Subscription) GetSuccess() uint {
	return e.Success
}

func (e *Subscription) GetIpAddress() string {
	return e.IpAddress
}

func (e *Subscription) GetIsRetry() bool {
	return e.IsRetry
}

func (e *Subscription) GetIsPurge() bool {
	return e.IsPurge
}

func (e *Subscription) GetIsActive() bool {
	return e.IsActive
}

func (e *Subscription) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *Subscription) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

func (s *Subscription) IsCreatedAtToday() bool {
	return s.CreatedAt.Format("2006-01-02") == time.Now().Format("2006-01-02")
}
