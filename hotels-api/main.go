package main

import (
	"hotels-api/router"
	"hotels-api/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
	// q "hotels-api/utils/queue"
)

var (
	ginRouter *gin.Engine
)

func main() {
	ginRouter = gin.Default()

	ginRouter.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Agrega los encabezados necesarios
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Si necesitas enviar credenciales (por ejemplo, cookies)
        
        if c.Request.Method == "OPTIONS" {
            // Respuesta a la solicitud OPTIONS
            c.JSON(200, nil)
            return
        }
        
        c.Next()
    })


	router.MapUrls(ginRouter)
	err := db.InitDB()
	defer db.DisconnectDB()
	// go q.QueueConnection()
	

	if err != nil {
		fmt.Println("Cannot init db")
		fmt.Println(err)
		return
	}
	fmt.Println("Starting server")
	ginRouter.Run(":8060")
}
