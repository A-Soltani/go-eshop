package domain

import "github.com/A-Soltani/go-eshop/src/services/ordering/domain/errors"

type OrderStatus int

const (
	OrderSubmitted OrderStatus = 1
	OrderPaid      OrderStatus = 2
)

type Order struct {
	Id         int
	Status     OrderStatus
	OrderItems []OrderItem
}

func AddOrder() *Order {
	o := &Order{}
	o.Status = OrderSubmitted
	return o
}

func AddOrderItem(productId int, units int, unitPrice float64, discount float64) (*OrderItem, error) {
	if productId < 1 {
		return nil, &errors.OrderingDomainError{Msg: "productId should be greater than 1"}
	}
	if float64(units)*unitPrice < discount {
		return nil, &errors.OrderingDomainError{Msg: "The total of order item is lower than applied discount"}
	}
	oi := &OrderItem{
		productId: productId,
		units:     units,
		unitPrice: unitPrice,
		discount:  discount}
	return oi, nil
}
