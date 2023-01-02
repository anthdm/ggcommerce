package main

import (
	"context"
	"fmt"

	"github.com/anthdm/ggcommerce/api"
	"github.com/anthdm/ggcommerce/store"

	"github.com/anthdm/weavebox"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleAPIError(ctx *weavebox.Context, err error) {
	fmt.Println(err)
}

func main() {
	app := weavebox.New()
	app.ErrorHandler = handleAPIError
	adminRoute := app.Box("/admin")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	productStore := store.NewMongoProductStore(client.Database("ggcommerce"))
	productHandler := api.NewProductHandler(productStore)

	adminRoute.Get("/product/:id", productHandler.HandleGetProductByID)
	adminRoute.Post("/product", productHandler.HandlePostProduct)

	app.Serve(3001)
}
