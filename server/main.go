package main

import (
	"os"
	"server/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func main() {

 	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	router.POST("/order/create", routes.AddOrder)

	router.GET("/waiter/:waiter", routes.GetOrdersByWaiter)
	router.GET("/orders", routes.GetOrders)
	router.GET("/order/:id/", routes.GetOrderById)

	router.PUT("/waiter/update/:id", routes.UpdateWaiter)
	router.PUT("/order/update/:id", routes.UpdateOrder)

	router.DELETE("/order/delete/:id", routes.DeleteOrder)

	router.Run(":" + port)

}