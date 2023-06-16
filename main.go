package main

import (
	"go_restfulapi/config"
	"go_restfulapi/controllers"
	"go_restfulapi/exception"
	"go_restfulapi/repository"
	"go_restfulapi/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	config.LoadEnv()
}

func main() {
	db := config.ConnectToDB()

	productRepository := repository.NewProductRepository()
	productService := services.NewProductService(db, productRepository)
	productController := controllers.NewProductController(productService)

	router := httprouter.New()
	router.GET("/api/v1/products", productController.FindAll)
	router.POST("/api/v1/products", productController.Create)
	router.PUT("/api/v1/products/:productId", productController.Update)
	router.GET("/api/v1/products/:productId", productController.FindById)
	router.DELETE("/api/v1/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	server.ListenAndServe()
}
