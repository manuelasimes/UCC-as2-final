package main

import (
	"user-res-api/app"
	"user-res-api/db"
	cache "user-res-api/cache"
	"fmt"
)

func main() {
	db.StartDbEngine()
	
    fmt.Println("Initializing cache")
	cache.Init_cache()
	
	app.StartRoute()
	
}