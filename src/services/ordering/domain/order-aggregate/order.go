package domain

type OrderStatus int

const (
	OrderSubmitted OrderStatus = 1
	OrderPaid      OrderStatus = 2
)

type Order struct {
	Id     int
	Status OrderStatus
}

func AddOrder() *Order {
	o := &Order{}
	o.Status = OrderSubmitted
	return o
}
