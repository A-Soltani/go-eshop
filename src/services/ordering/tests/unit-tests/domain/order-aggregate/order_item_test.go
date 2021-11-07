package domain_test

import (
	"testing"

	errors "github.com/A-Soltani/go-eshop/src/services/ordering/domain/errors"
	domain "github.com/A-Soltani/go-eshop/src/services/ordering/domain/order-aggregate"
	"github.com/stretchr/testify/assert"
)

func TestAddOrderItem_InvalidUnits_ErrorShouldBeReturned(t *testing.T) {
	productId := 1
	units := -1

	orderItem, _ := domain.AddOrderItem(productId)
	err := orderItem.AddUnits(units)

	assert.IsType(t, &errors.OrderingDomainError{}, err)
}
