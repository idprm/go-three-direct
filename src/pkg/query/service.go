package query

import (
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func GetServiceById(id int) (model.Service, error) {
	var s model.Service
	sql := "SELECT id, code, name, auth_user, auth_pass, day, charge, trial_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE id = ? AND is_active = true LIMIT 1"
	db := database.Datasource.SqlDB()
	err := db.QueryRow(sql, id).Scan(&s.ID, &s.Code, &s.Name, &s.AuthUser, &s.AuthPass, &s.Day, &s.Charge, &s.TrialDay, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return s, err
	}
	return s, nil
}

func GetServiceByCode(code string) (model.Service, error) {
	var s model.Service
	sql := "SELECT id, code, name, auth_user, auth_pass, day, charge, trial_day, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE code = ? AND is_active = true LIMIT 1"
	db := database.Datasource.SqlDB()
	err := db.QueryRow(sql, code).Scan(&s.ID, &s.Code, &s.Name, &s.AuthUser, &s.AuthPass, &s.Day, &s.Charge, &s.TrialDay, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return s, err
	}
	return s, nil
}
