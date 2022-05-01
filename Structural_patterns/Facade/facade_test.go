package t_test

import (
	"DESIGN_PATTERNS/Structural_patterns/Facade/facade"
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	walletFacade := facade.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		t.Error(err)
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		t.Error(err)
	}
}
