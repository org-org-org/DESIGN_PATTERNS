package observer

type Subject interface {
	Register(Observer observer)
	Deregister(Observer observer)
	NotifyAll()
}
