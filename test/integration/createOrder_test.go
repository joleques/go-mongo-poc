package integration

import (
	"github.com/joleques/go-mongo-poc/src/useCase"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestCreateOrderLoadSuccessfully(t *testing.T) {
	var products []string
	for i := 0; i < 48; i++ {
		element := "test" + strconv.Itoa(i)
		products = append(products, element)
	}
	totalRequests := 40
	wg.Add(totalRequests)
	for j := 0; j < totalRequests; j++ {
		identifier := "identifier" + strconv.Itoa(j)
		go useCase.CreateOrder(identifier, products, &wg)
	}
	wg.Wait()
}

func TestCreateOrderSuccessfully(t *testing.T) {
	var products []string
	for i := 0; i < 48; i++ {
		element := "test" + strconv.Itoa(i)
		products = append(products, element)
	}
	wg.Add(1)
	identifier := "identifier"
	err := useCase.CreateOrder(identifier, products, &wg)
	assert.Nil(t, err)
}

func TestCreateOrderWithErrorWhenIdentifierIsEmpty(t *testing.T) {
	identifier := ""
	products := []string{"test_1", "test_2", "test_3", "test_4", "test_5"}
	err := useCase.CreateOrder(identifier, products, &wg)
	assert.Equal(t, "error: identifier empty", err.Error())
}

func TestCreateOrderWithErrorWhenProductsIsEmpty(t *testing.T) {
	identifier := "12_test"
	products := []string{}
	err := useCase.CreateOrder(identifier, products, &wg)
	assert.Equal(t, "error: products empty", err.Error())
}

func TestCreateOrderWithErrorWhenProductsBiggerThen5(t *testing.T) {
	identifier := "12_test"
	var products []string
	for i := 0; i < 55; i++ {
		element := "test" + strconv.Itoa(i)
		products = append(products, element)
	}
	err := useCase.CreateOrder(identifier, products, &wg)
	assert.Equal(t, "error: Product list cannot be longer than 50", err.Error())
}
