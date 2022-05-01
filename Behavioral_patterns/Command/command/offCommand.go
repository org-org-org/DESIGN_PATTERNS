package command

type OffCommand struct {
	Device device
}

func (c *OffCommand) execute() {
	c.Device.off()
}
