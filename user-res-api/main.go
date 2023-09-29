package main

import (
	"UCC-as2-final/app"
	"UCC-as2-final/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}