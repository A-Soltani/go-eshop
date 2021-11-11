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

	o := domain.AddOrder()
	_ = o.AddOrderItem(productId, units, unitPrice, discount)
	orderItem := o.OrderItems[0]
	err := orderItem.AddUnits(newUnits)

	assert.IsType(t, &errors.OrderingDomainError{}, err)
}

func TestAddOrderItem_InvalidDiscount_ErrorShouldBeReturned(t *testing.T) {
	productId := 1
	units := 1
	unitPrice := 1.5
	discount := 30.5

	o := domain.AddOrder()
	err := o.AddOrderItem(productId, units, unitPrice, discount)

	assert.IsType(t, &errors.OrderingDomainError{}, err)
}
