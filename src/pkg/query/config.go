package query

import (
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func GetConfig(name string) (model.Config, error) {
	var config model.Config
	sqlStatement := "SELECT value FROM configs WHERE name = ? LIMIT 1"
	db := database.Datasource.SqlDB()
	err := db.QueryRow(sqlStatement, name).Scan(&config.Value)
	if err != nil {
		return config, err
	}
	return config, nil
}
