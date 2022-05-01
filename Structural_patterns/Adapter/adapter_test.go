package t_test

import (
	"DESIGN_PATTERNS/Structural_patterns/Adapter/adapter"
	"testing"
)

func TestAdapter(t *testing.T) {
	client := &adapter.Client{}
	mac := &adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
