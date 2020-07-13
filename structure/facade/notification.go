package facade

import "fmt"

type notification struct{}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Send wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Send wallet debit notification")
}
