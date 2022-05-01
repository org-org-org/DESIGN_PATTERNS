package bridge

import "fmt"

type Mac struct {
	printer printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *Mac) SetPrinter(p printer) {
	m.printer = p
}
