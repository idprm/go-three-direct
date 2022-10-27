package query

import (
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func GetContent(serviceId int, name string) (model.Content, error) {
	var content model.Content
	sqlStatement := "SELECT value FROM contents WHERE service_id = ? AND name = ? LIMIT 1"
	db := database.Datasource.SqlDB()
	err := db.QueryRow(sqlStatement, serviceId, name).Scan(&content.Value)
	if err != nil {
		return content, err
	}
	return content, nil
}
