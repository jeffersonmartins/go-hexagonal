package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/jeffersonmartins/go-hexagonal/application"
	"github.com/google/uuid"
)

func TestProduct_Enable(t *testing.T) {    
    product := application.Product{}
    product.Name = "Product 1"
    product.Status = "disabled"
    product.Price = 10
    err := product.Enable()
    
    require.Nil(t, err)

    product.Price = -10
    err = product.Enable()
    require.EqualError(t, err, "the price must be greater than zero to enable the product")
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.EqualError(t, err, "the price must be zero in order to disable the product")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10
	
	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.EqualError(t, err, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.EqualError(t, err, "the price must be greater or equal zero", err.Error())
}