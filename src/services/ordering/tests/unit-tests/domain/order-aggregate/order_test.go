package domain_test

import (
	"testing"

	domain "github.com/A-Soltani/go-eshop/src/services/ordering/domain/order-aggregate"
	"github.com/stretchr/testify/assert"
)

func TestAddOrder_ValidStatus_ValidOrderShouldBeReturned(t *testing.T) {
	o := domain.AddOrder()

	assert.True(t, o.Status == domain.OrderSubmitted)
}

func TestAddOrderItem_NewOrderItem_NewOrderItemShouldBeAddedToOrderItems(t *testing.T) {
	productId := 1
	units := 1
	unitPrice := 1.5
	discount := 0.5

	o := domain.AddOrder()
	_ = o.AddOrderItem(productId, units, unitPrice, discount)

	assert.Equal(t, len(o.OrderItems), 1)
}

func TestCancelOrder_SetCancelledStatus_OrderCancelledDomainEventShouldBeAdded(t *testing.T) {

	o := domain.AddOrder()
	o.Cancel()

	assert.True(t, o.Status == domain.OrderCancelled)
}
