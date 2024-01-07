package application_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sandro-dev/golang-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.New().String()
	product.Name = "Product A"
	product.Status = application.DISABLED
	product.Price = 100

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable product", err.Error())

}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.New().String()
	product.Name = "Product B"
	product.Status = application.DISABLED
	product.Price = 100

	err := product.Enable()
	require.Nil(t, err)

	err = product.Disable()
	require.Equal(t, "the price must be equals to zero", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)

}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.New().String()
	product.Name = "Product C"
	product.Status = application.ENABLED
	product.Price = 100

	product.Price = -10

	_, err := product.IsValid()
	require.Equal(t, "the price must be greater than zero", err.Error())

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "status must be ENABLED or DISABLED", err.Error())

	product.Price = 100
	err = product.Enable()
	require.Nil(t, err)

}
