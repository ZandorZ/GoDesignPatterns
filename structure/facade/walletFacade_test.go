package facade

import (
	"fmt"
	"log"
	"testing"
)

func TestWalletFacade(t *testing.T) {
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
