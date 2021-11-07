package errors

type OrderingDomainError struct {
	Msg string
}

func (ode *OrderingDomainError) Error() string {
	return ode.Msg
}
