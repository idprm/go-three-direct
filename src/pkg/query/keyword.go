package query

import (
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func GetKeyword(name string) (model.Keyword, error) {
	var keyword model.Keyword
	sqlStatement := "SELECT id, name FROM keywords WHERE name = ? LIMIT 1"
	db := database.Datasource.SqlDB()
	err := db.QueryRow(sqlStatement, name).Scan(&keyword.ID, &keyword.Name)
	if err != nil {
		return keyword, err
	}
	return keyword, nil
}
