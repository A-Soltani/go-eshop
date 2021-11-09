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
