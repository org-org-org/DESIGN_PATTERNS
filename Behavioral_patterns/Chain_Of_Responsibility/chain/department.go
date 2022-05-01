package chain

type department interface {
	execute(*Patient)
	SetNext(department)
}
