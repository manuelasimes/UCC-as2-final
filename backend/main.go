package main

import (
	"backend/app"
	"backend/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
