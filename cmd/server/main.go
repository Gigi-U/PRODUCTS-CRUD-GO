package main

import (
	"log"
	"time"

	handlerPing "github.com/Gigi-U/PRODUCTS-CRUD-GO.git/cmd/server/handler/ping"
	handlerProducto "github.com/Gigi-U/PRODUCTS-CRUD-GO.git/cmd/server/handler/products"
	
	"github.com/Gigi-U/PRODUCTS-CRUD-GO.git/internal/models"
	"github.com/Gigi-U/PRODUCTS-CRUD-GO.git/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {

	// Carga la base de datos en memoria
	db := LoadStore()

	// Ping.
	controllerPing := handlerPing.NewControllerPing()

	// Products.
	repostory := products.NewMemoryRepository(db)
	service := products.NewServiceProduct(repostory)
	controllerProduct := handlerProducto.NewControllerProducts(service)

	engine := gin.Default()

	group := engine.Group("/api/v1")
	{
		group.GET("/ping", controllerPing.HandlerPing())

		grupoProducto := group.Group("/producto")
		{
			grupoProducto.POST("", controllerProduct.HandlerCreate())
			grupoProducto.GET("", controllerProduct.HandlerGetAll())
			grupoProducto.GET("/:id", controllerProduct.HandlerGetByID())
			grupoProducto.PUT("/:id", controllerProduct.HandlerUpdate())
			grupoProducto.DELETE("/:id", controllerProduct.HandlerDelete())

		}

	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// LoadStore carga la base de datos en memoria
func LoadStore() []models.Producto {
	return []models.Producto{
		{
			Id:          "1",
			Name:        "Coco Cola",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.5,
		},
		{
			Id:          "2",
			Name:        "Pepsito",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.5,
		},
		{
			Id:          "3",
			Name:        "Fantastica",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.5,
		},
	}
}

// clase 11