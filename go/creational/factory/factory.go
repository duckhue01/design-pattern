package factory

import (
	"fmt"
)

type (
	PaymentMethod interface {
		Pay(amount int) string
	}
)

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {

	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)

	}
}

type CashPM struct{}

type DebitCardPM struct{}

func (cp *CashPM) Pay(amount int) string {
	return "cash payment method"
}
func (dcp *DebitCardPM) Pay(amount int) string {
	return "debit card payment method"
}
