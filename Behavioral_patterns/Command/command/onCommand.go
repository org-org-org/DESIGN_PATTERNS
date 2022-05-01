package command

type OnCommand struct {
	Device device
}

func (c *OnCommand) execute() {
	c.Device.on()
}
