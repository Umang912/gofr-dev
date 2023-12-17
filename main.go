package main

import (
	"GO-LANG/datastore"
	"GO-LANG/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	// Assuming `datastore.New()` initializes an Inventory datastore instance
	s := datastore.NewInventory()
	h := handler.New(*s)

	app.GET("/product/{ID}", h.GetProductByID)
	app.POST("/product", h.AddProduct)
	app.DELETE("/product/{ID}", h.RemoveProduct)
	app.PUT("/product/{ID}", h.UpdateProduct)

	app.Server.HTTP.Port = 9092
	app.Start()
}
