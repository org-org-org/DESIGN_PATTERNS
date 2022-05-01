package t_test

import (
	"DESIGN_PATTERNS/Behavioral_patterns/Iterator/iterator"
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	user1 := &iterator.User{
		Name: "a",
		Age:  30,
	}
	user2 := &iterator.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}

	it := userCollection.CreateIterator()

	for it.HasNext() {
		user := it.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
