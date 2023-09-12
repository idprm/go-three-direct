package telco

import "waki.mobi/go-yatta-h3i/src/config"

type Telco struct {
	cfg *config.Secret
}

func NewTelco(cfg *config.Secret) *Telco {
	return &Telco{
		cfg: cfg,
	}
}
