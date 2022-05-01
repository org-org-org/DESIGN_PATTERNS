package chain

import "fmt"

type Cashier struct {
	next department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
		return
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next department) {
	c.next = next
}
