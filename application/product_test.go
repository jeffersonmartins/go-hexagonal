package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
)
func TestProduct_Enable(t *testing.T) {	

	product := application.Product{}
	product.Name = "Product 1"
	product.Status = "disabled"
	product.Price = 10
	err := product.Enable()
	
	require.Nil(t, err)	;

}