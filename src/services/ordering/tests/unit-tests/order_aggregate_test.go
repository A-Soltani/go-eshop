package domain_test

import (
	"testing"

	"github.com/A-Soltani/go-eshop/src/services/ordering/domain"
	"github.com/stretchr/testify/assert"
)

func TestAddOrder(t *testing.T) {
	productId := 1
	units := -1

	orderItem, _ := domain.AddOrderItem(productId)
	err := orderItem.AddUnits(units)

	assert.EqualError(t, err, "units can't be negative")
}
