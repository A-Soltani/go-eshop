package domain

import "errors"

type OrderItem struct {
	productId int
	units     int
}

func AddOrderItem(prodductId int) (*OrderItem, error) {
	if prodductId < 1 {
		return nil, errors.New("productId should be greater than 1")
	}
	oi := &OrderItem{productId: prodductId}
	return oi, nil
}

func (oi *OrderItem) AddUnits(units int) error {
	if units < 0 {
		return errors.New("units can't be negative")
	}

	oi.units += units
	return nil
}
