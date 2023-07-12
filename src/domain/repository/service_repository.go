package repository

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/domain/entity"
)

const (
	queryServiceById   = "SELECT id, code, name, auth_user, auth_pass, day, charge, purge_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE id = ? AND is_active = true LIMIT 1"
	queryServiceByCode = "SELECT id, code, name, auth_user, auth_pass, day, charge, purge_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE code = ? AND is_active = true LIMIT 1"
	queryServiceByName = "SELECT id, code, name, auth_user, auth_pass, day, charge, purge_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE name = ? AND is_active = true LIMIT 1"
)

type ServiceRepository struct {
	db *sql.DB
}

type IServiceRepository interface {
	GetServiceById(int) (*entity.Service, error)
	GetServiceByCode(string) (*entity.Service, error)
	GetServiceByName(string) (*entity.Service, error)
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (r *ServiceRepository) GetServiceById(id int) (*entity.Service, error) {
	var s entity.Service
	err := r.db.QueryRow(queryServiceById, id).Scan(&s.ID, &s.Code, &s.Name, &s.AuthUser, &s.AuthPass, &s.Day, &s.Charge, &s.PurgeDay, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *ServiceRepository) GetServiceByCode(code string) (*entity.Service, error) {
	var s entity.Service
	err := r.db.QueryRow(queryServiceByCode, code).Scan(&s.ID, &s.Code, &s.Name, &s.AuthUser, &s.AuthPass, &s.Day, &s.Charge, &s.PurgeDay, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *ServiceRepository) GetServiceByName(name string) (*entity.Service, error) {
	var s entity.Service
	err := r.db.QueryRow(queryServiceByName, name).Scan(&s.ID, &s.Code, &s.Name, &s.AuthUser, &s.AuthPass, &s.Day, &s.Charge, &s.PurgeDay, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
