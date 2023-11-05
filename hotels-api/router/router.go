package router

import (
	hotelController "hotels-api/controllers/hotel"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func MapUrls(router *gin.Engine) {

	// router.Use(corsMiddleware())
	// Products Mapping
	router.GET("/hotels/:id", hotelController.Get)
	router.POST("/hotels", hotelController.Insert)
	router.PUT("/hotels/:id", hotelController.Update)

	fmt.Println("Finishing mappings configurations")
}

func corsMiddleware() gin.HandlerFunc {
    // Define your CORS configuration here
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your frontend's origin
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
    config.AllowHeaders = []string{"Content-Type", "Authorization"}

    return cors.New(config)
}
