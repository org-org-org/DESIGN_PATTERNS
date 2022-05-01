package command

type command interface {
	execute()
}

type Button struct {
	Command command
}

func (b *Button) Press() {
	b.Command.execute()
}
