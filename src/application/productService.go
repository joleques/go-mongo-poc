package application

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	Identifier string             `bson:"identifier"`
	Products   []string           `bson:"products"`
}

func NewOrder(identifier string, products []string) (*Order, error) {
	orderNew := Order{ID: primitive.NewObjectID(), Identifier: identifier, Products: products, UpdatedAt: time.Now(), CreatedAt: time.Now()}
	err := orderNew.Verify()
	if err != nil {
		return nil, err
	}
	return &orderNew, nil
}

func (order Order) Verify() error {
	if order.Identifier == "" {
		return errors.New("error: identifier empty")
	}

	if order.Products == nil {
		return errors.New("error: products nil")
	}

	if len(order.Products) == 0 {
		return errors.New("error: products empty")
	}
	if len(order.Products) > 50 {
		return errors.New("error: Product list cannot be longer than 50")
	}

	return nil
}

func (order Order) TotalProducts() int {
	return len(order.Products)
}
