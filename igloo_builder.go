package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floors     int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) SetNumberOfFloors() {
	b.floors = 1
}

func (b *IglooBuilder) GetHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floors:     b.floors,
	}
}
