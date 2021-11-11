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

func (o *Order) AddOrderItem(productId int, units int, unitPrice float64, discount float64) error {
	if productId < 1 {
		return &errors.OrderingDomainError{Msg: "productId should be greater than 1"}
	}
	if float64(units)*unitPrice < discount {
		return &errors.OrderingDomainError{Msg: "The total of order item is lower than applied discount"}
	}

	oi := OrderItem{
		productId: productId,
		units:     units,
		unitPrice: unitPrice,
		discount:  discount}

	o.OrderItems = append(o.OrderItems, oi)
	return nil
}
