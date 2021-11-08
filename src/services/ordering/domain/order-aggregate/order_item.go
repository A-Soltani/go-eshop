package domain

import "github.com/A-Soltani/go-eshop/src/services/ordering/domain/errors"

type OrderItem struct {
	productId int
	units     int
	unitPrice float64
	discount  float64
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

func (oi *OrderItem) AddUnits(units int) error {
	if units < 0 {
		return &errors.OrderingDomainError{Msg: "units can't be negative"}
	}

	oi.units += units
	return nil
}
