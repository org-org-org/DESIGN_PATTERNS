package bridge

import "fmt"

type Hp struct {
}

func (p *Hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}
