package main

import (
	"waki.mobi/go-yatta-h3i/src/cmd"
	"waki.mobi/go-yatta-h3i/src/database"
)

func init() {
	// Setup database
	database.Connect()
}

func main() {
	// Setup cobra
	cmd.Execute()

}
