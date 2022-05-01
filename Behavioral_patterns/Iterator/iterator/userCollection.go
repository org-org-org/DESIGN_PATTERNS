package iterator

type Iterator interface {
	HasNext() bool
	GetNext() *User
}

type Collection interface {
	createIterator() Iterator
}

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &userIterator{
		users: u.Users,
	}
}
