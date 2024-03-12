package route

import (
	"assignment2/controllers"
	"assignment2/database"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	orderController := controllers.OrderController{
		DB: database.DB,
	}

	orderRoute := r.Group("/orders")
	orderRoute.GET("/", orderController.GetAll)
	orderRoute.POST("/", orderController.Create)
	orderRoute.PUT("/:orderId", orderController.Update)
	orderRoute.DELETE("/:orderId", orderController.Delete)
}
