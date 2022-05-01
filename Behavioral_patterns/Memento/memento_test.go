package t_test

import (
	"DESIGN_PATTERNS/Behavioral_patterns/Memento/memento"
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker := &memento.Caretaker{
		MementoArray: make([]*memento.Memento, 0),
	}

	originator := &memento.Originator{
		State: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())
}
