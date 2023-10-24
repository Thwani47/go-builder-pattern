package main

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumberOfFloors()
	GetHouse() House
}

func GetBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return NewNormalBuilder()
	case "igloo":
		return NewIglooBuilder()
	default:
		return nil
	}
}
