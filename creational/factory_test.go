package creational

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {

	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "payed using cash") {
		t.Error("Error in cash payment method!")
	}

	t.Log(msg)

}

func TestGetPaymentMethodDebitCard(t *testing.T) {

	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	msg := payment.Pay(20.44)
	if !strings.Contains(msg, "payed using debit card") {
		t.Error("Error in cash payment method!")
	}

	t.Log(msg)

}

func TestGetPaymentMethodUnknown(t *testing.T) {

	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment method with ID 20 must return an error!")
	}
	t.Log(err)
}
