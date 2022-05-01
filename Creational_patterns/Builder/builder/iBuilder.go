package builder

type house struct {
	WindowType string
	DoorType   string
	Floor      int
}

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
