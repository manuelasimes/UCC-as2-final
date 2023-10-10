package main

import (
	"UCC-as2-final/app"
	q "UCC-as2-final/utils/connections"
)

func main() {
	go q.QueueConnection()
	app.StartRoute()
}