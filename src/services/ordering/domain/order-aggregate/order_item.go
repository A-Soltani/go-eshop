package domain

import "github.com/A-Soltani/go-eshop/src/services/ordering/domain/errors"

type OrderItem struct {
	productId int
	units     int
}

func AddOrderItem(prodductId int) (*OrderItem, error) {
	if prodductId < 1 {
		// return nil, errors.New("productId should be greater than 1")
		return nil, &errors.OrderingDomainError{Msg: "productId should be greater than 1"}
	}
	oi := &OrderItem{productId: prodductId}
	return oi, nil
}

func (oi *OrderItem) AddUnits(units int) error {
	if units < 0 {
		return &errors.OrderingDomainError{Msg: "units can't be negative"}
	}

	oi.units += units
	return nil
}
