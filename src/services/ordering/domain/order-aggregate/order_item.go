package domain

import "github.com/A-Soltani/go-eshop/src/services/ordering/domain/errors"

type OrderItem struct {
	productId int
	units     int
	unitPrice float64
	discount  float64
}

func (oi *OrderItem) AddUnits(units int) error {
	if units < 0 {
		return &errors.OrderingDomainError{Msg: "units can't be negative"}
	}

	oi.units += units
	return nil
}
