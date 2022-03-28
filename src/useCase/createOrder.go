package useCase

import (
	"context"
	"github.com/joleques/go-mongo-poc/src/application"
	mongo2 "github.com/joleques/go-mongo-poc/src/infra/mongo"
	"sync"
)

var ctx = context.TODO()

func CreateOrder(identifier string, products []string, wg *sync.WaitGroup) error {
	order, err := application.NewOrder(identifier, products)
	if err != nil {
		return err
	}
	client := mongo2.NewDataBase()
	client.InsertOrder(ctx, order)
	wg.Done()
	return nil
}
