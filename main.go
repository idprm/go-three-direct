package main

import (
	"waki.mobi/go-yatta-h3i/src/cmd"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

func init() {
	// Setup database
	database.Connect()

	// Setup logging
	util.WriteLog()
}

func main() {
	// Setup cobra
	cmd.Execute()
}
