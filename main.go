package main

import (
	"waki.mobi/go-yatta-h3i/src/cmd"
	"waki.mobi/go-yatta-h3i/src/database"
)

func main() {
	// Setup database
	database.Connect()
	// Setup cobra
	cmd.Execute()
}
