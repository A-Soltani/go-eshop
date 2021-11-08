package domain_test

import (
	"testing"

	errors "github.com/A-Soltani/go-eshop/src/services/ordering/domain/errors"
	domain "github.com/A-Soltani/go-eshop/src/services/ordering/domain/order-aggregate"
	"github.com/stretchr/testify/assert"
)

func TestAddOrderItem_InvalidUnits_ErrorShouldBeReturned(t *testing.T) {
	productId := 1
	units := 1
	newUnits := -1
	unitPrice := 1.5
	discount := 0.5

	orderItem, _ := domain.AddOrderItem(productId, units, unitPrice, discount)
	err := orderItem.AddUnits(newUnits)

	assert.IsType(t, &errors.OrderingDomainError{}, err)
}

func TestAddOrderItem_InvalidDiscount_ErrorShouldBeReturned(t *testing.T) {
	productId := 1
	units := 1
	unitPrice := 1.5
	discount := 30.5

	_, err := domain.AddOrderItem(productId, units, unitPrice, discount)

	assert.IsType(t, &errors.OrderingDomainError{}, err)
}
