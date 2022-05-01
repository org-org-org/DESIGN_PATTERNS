package bridge

import "fmt"

type Windows struct {
	printer printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *Windows) SetPrinter(p printer) {
	w.printer = p
}
