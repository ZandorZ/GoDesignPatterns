package creational

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

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
		return nil, errors.New(fmt.Sprintf("Invalid payment %d method!", m))
	}

}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f USD payed using cash", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f USD payed using debit card", amount)
}
