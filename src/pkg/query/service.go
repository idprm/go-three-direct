package query

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

type ServiceRepository struct {
	db *sql.DB
}

type IServiceRepository interface {
	GetServiceById(int) (model.Service, error)
	GetServiceByCode(string) (model.Service, error)
	GetServiceByName(string) (model.Service, error)
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (r *ServiceRepository) GetServiceById(id int) (model.Service, error) {
	var s model.Service
	sql := "SELECT id, code, name, auth_user, auth_pass, day, charge, purge_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE id = ? AND is_active = true LIMIT 1"
	err := r.db.QueryRow(sql, id).Scan(&s.ID, &s.Code, &s.Name, &s.AuthUser, &s.AuthPass, &s.Day, &s.Charge, &s.PurgeDay, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *ServiceRepository) GetServiceByCode(code string) (model.Service, error) {
	var s model.Service
	sql := "SELECT id, code, name, auth_user, auth_pass, day, charge, purge_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE code = ? AND is_active = true LIMIT 1"
	err := r.db.QueryRow(sql, code).Scan(&s.ID, &s.Code, &s.Name, &s.AuthUser, &s.AuthPass, &s.Day, &s.Charge, &s.PurgeDay, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *ServiceRepository) GetServiceByName(name string) (model.Service, error) {
	// SELECT id, code, name, auth_user, auth_pass, day, charge, purge_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM yatta_db.services WHERE name = 'GM' AND is_active = true LIMIT 1;
	var s model.Service
	sql := "SELECT id, code, name, auth_user, auth_pass, day, charge, purge_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE name = ? AND is_active = true LIMIT 1"
	err := r.db.QueryRow(sql, name).Scan(&s.ID, &s.Code, &s.Name, &s.AuthUser, &s.AuthPass, &s.Day, &s.Charge, &s.PurgeDay, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return s, err
	}
	return s, nil
}
