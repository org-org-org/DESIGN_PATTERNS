package State

import (
	"DESIGN_PATTERNS/Behavioral_patterns/State/state"
	"testing"
)

func TestStrategy(t *testing.T) {
	vendingMachine := state.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		t.Error(err)
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		t.Error(err)
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		t.Error(err)
	}

	err = vendingMachine.AddItem(2)
	if err != nil {
		t.Error(err)
	}

	err = vendingMachine.RequestItem()
	if err != nil {
		t.Error(err)
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		t.Error(err)
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		t.Error(err)
	}
}
