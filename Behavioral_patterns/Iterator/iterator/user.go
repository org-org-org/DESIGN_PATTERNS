package iterator

type User struct {
	Name string
	Age  int
}

type userIterator struct {
	index int
	users []*User
}

func (u *userIterator) HasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}
func (u *userIterator) GetNext() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
