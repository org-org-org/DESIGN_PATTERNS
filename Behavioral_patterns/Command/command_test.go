package t_test

import (
	"DESIGN_PATTERNS/Behavioral_patterns/Command/command"
	"testing"
)

func TestCommand(t *testing.T) {
	tv := &command.Tv{}

	onCommand := &command.OnCommand{
		Device: tv,
	}

	offCommand := &command.OffCommand{
		Device: tv,
	}

	onButton := &command.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &command.Button{
		Command: offCommand,
	}
	offButton.Press()
}
