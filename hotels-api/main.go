package main

import (
	"hotels-api/router"
	"hotels-api/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	ginRouter *gin.Engine
)

func main() {
	ginRouter = gin.Default()
	router.MapUrls(ginRouter)
	err := db.InitDB()
	defer db.DisconnectDB()

	if err != nil {
		fmt.Println("Cannot init db")
		fmt.Println(err)
		return
	}
	fmt.Println("Starting server")
	ginRouter.Run(":8090")
}
