package unit

import (
	"github.com/joleques/go-mongo-poc/src/application"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_orderSuccessfully(t *testing.T) {
	order, err := application.NewOrder("test", []string{"test1", "test2"})
	assert.Nil(t, err)
	assert.Equal(t, "test", order.Identifier)
	assert.Equal(t, 2, len(order.Products))

	products := order.Products
	assert.Equal(t, "test1", products[0])
	assert.Equal(t, "test2", products[1])
}

func Test_orderErrorWhenIdentifierEmpty(t *testing.T) {
	order, err := application.NewOrder("", []string{"test1", "test2"})
	assert.Nil(t, order)
	assert.Equal(t, "error: identifier empty", err.Error())
}

func Test_orderErrorWhenProductsEmpty(t *testing.T) {
	order, err := application.NewOrder("test", []string{})
	assert.Nil(t, order)
	assert.Equal(t, "error: products empty", err.Error())
}

func Test_orderErrorWhenProductsNil(t *testing.T) {
	order, err := application.NewOrder("test", nil)
	assert.Nil(t, order)
	assert.Equal(t, "error: products nil", err.Error())
}

func Test_orderErrorWhenProductsBiggerThen5(t *testing.T) {
	var products []string
	for i := 0; i < 55; i++ {
		element := "test" + strconv.Itoa(i)
		products = append(products, element)
	}
	order, err := application.NewOrder("test", products)
	assert.Nil(t, order)
	assert.Equal(t, "error: Product list cannot be longer than 50", err.Error())
}
