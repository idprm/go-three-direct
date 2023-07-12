package entity

import "time"

type Blacklist struct {
	ID        uint64    `json:"id"`
	Msisdn    string    `json:"msisdn"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *Blacklist) GetId() uint64 {
	return e.ID
}

func (e *Blacklist) GetMsisdn() string {
	return e.Msisdn
}
