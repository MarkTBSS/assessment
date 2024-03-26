package main

import (
	"github.com/MarkTBSS/assessment/configs"
	"github.com/MarkTBSS/assessment/databases"
)

func main() {
	configs.LoadConfig()
	databases.ConnectDatabase()
}
