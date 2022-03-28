package mongo

import (
	"context"
	"github.com/joleques/go-mongo-poc/src/application"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection
var ctx = context.TODO()

type MongoDB struct {
	dataBase *mongo.Database
}

func NewDataBase() MongoDB {
	clientOptions := options.Client().ApplyURI("mongodb://golang-test:test123@localhost:27017/scheduler?authSource=golang-test&readPreference=primary&appname=MongoDB%20Compass&ssl=false")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return MongoDB{dataBase: client.Database("golang-test")}
}

func (client MongoDB) InsertOrder(ctx context.Context, order *application.Order) {
	client.dataBase.Collection("order").InsertOne(ctx, order)
}
