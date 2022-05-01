package t_test

import (
	"DESIGN_PATTERNS/Behavioral_patterns/Observer/observer"
	"testing"
)

func TestObserver(t *testing.T) {

	shirtItem := observer.NewItem("Nike Shirt")

	observerFirst := &observer.Customer{Id: "abc@gmail.com"}
	observerSecond := &observer.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.NotifyAll()
}
