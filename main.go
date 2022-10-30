package main

import (
	"log"

	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

func init() {
	// Setup database
	// database.Connect()
}

func main() {
	// Setup cobra
	// cmd.Execute()

	test, yes := util.KeywordDefine("REG KERENYT")
	log.Println(test)
	log.Println(yes)
}
