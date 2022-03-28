package main

import (
	"context"
	"fmt"
	"github.com/joleques/go-mongo-poc/src/application"
	mongo2 "github.com/joleques/go-mongo-poc/src/infra/mongo"
)

var ctx = context.TODO()

func main() {

	client := mongo2.NewDataBase()

	order, err := application.NewOrder("teste", []string{"test1", "test2"})
	if err != nil {
		panic(err)
	}
	client.InsertOrder(ctx, order)

	fmt.Println("Teste criado no Mongo....")
}
