package query

import (
	"waki.mobi/go-yatta-h3i/src/database"
)

func GetCountBlacklist(msisdn string) (int, error) {
	var count int
	sqlStatement := "SELECT COUNT(*) as count FROM blacklists WHERE msisdn = ? LIMIT 1"
	db := database.Datasource.SqlDB()
	err := db.QueryRow(sqlStatement, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}
