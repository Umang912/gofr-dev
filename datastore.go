package datastore

import (
	"GO-LANG/model"
	"os"
	"strconv"

	"context"
	"fmt"
	"log"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Inventory struct {
	client *mongo.Client
}

func NewInventory() *Inventory {
	return &Inventory{}
}

func (s *Inventory) connectMongoDB(ctx *gofr.Context) *mongo.Collection {
	if s.client == nil {
		uri := "mongodb+srv://prakharjain496:golangdev@golangdb.v2pxtfw.mongodb.net/?retryWrites=true&w=majority"
		if uri == "" {
			log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		}
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		s.client = client
	}
	fmt.Println("Connected to MongoDB")
	return s.client.Database("sample_inventory").Collection("Products")
}

func (s *Inventory) GetByID(ctx *gofr.Context, ID string) (*model.Product, error) {
	coll := s.connectMongoDB(ctx)
	var result model.Product
	i, errr := strconv.Atoi(ID)
	if errr != nil {
		return nil, errr
	}
	fmt.Println("Product ID:", i)

	err := coll.FindOne(context.Background(), bson.M{"id": i}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne ERROR:", err)
		return nil, err
	}

	fmt.Println(coll.FindOne(context.Background(), bson.M{"id": i}))
	return &result, nil
}

func (s *Inventory) AddProduct(ctx *gofr.Context, product *model.Product) (*model.Product, error) {
	coll := s.connectMongoDB(ctx)

	data := model.Product{ID: product.ID, Name: product.Name, Description: product.Description, Price: product.Price}
	result, insertErr := coll.InsertOne(ctx, data)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("Data inserted with objectID: ", result.InsertedID)
	}

	return &data, nil
}

func (s *Inventory) RemoveProduct(ctx *gofr.Context, id int) error {
	coll := s.connectMongoDB(ctx)

	filter := bson.M{"id": id}
	result, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of products deleted: %d\n", result.DeletedCount)
	return nil
}

func (s *Inventory) UpdateProduct(ctx *gofr.Context, product *model.Product) (*model.Product, error) {
	coll := s.connectMongoDB(ctx)

	existingProduct, err := s.GetByID(ctx, strconv.Itoa(product.ID))
	if err != nil {
		fmt.Println("Error fetching existing product data:", err)
		return nil, err
	}

	updatedProduct := mergeProducts(existingProduct, product)

	filter := bson.M{"id": product.ID}
	update := bson.M{"$set": updatedProduct}

	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Failed to update product")
		return nil, errors.DB{Err: err}
	}

	fmt.Println("Product updated successfully", product.ID)
	return updatedProduct, nil
}

func mergeProducts(existing *model.Product, update *model.Product) *model.Product {
	if update.Name != "" {
		existing.Name = update.Name
	}
	if update.Description != "" {
		existing.Description = update.Description
	}
	if update.Price != 0 {
		existing.Price = update.Price
	}
	return existing
}
