package adapter

import "fmt"

type WindowsAdapter struct {
	WindowMachine *Windows
}

func (w *WindowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.insertIntoUSBPort()
}
