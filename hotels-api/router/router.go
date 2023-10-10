package router

import (
	hotelController "hotels-api/controllers/hotel"
	"fmt"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/hotels/:id", hotelController.Get)
	router.POST("/hotels", hotelController.Insert)

	fmt.Println("Finishing mappings configurations")
}
