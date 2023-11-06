package main

import (
	"search-api/app"
	q "search-api/utils/connections"
)

func main() {
	go q.QueueConnection()
	app.StartRoute()
}