package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floors     int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) SetNumberOfFloors() {
	b.floors = 2
}

func (b *NormalBuilder) GetHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floors:     b.floors,
	}
}
