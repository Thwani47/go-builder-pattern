package main

import "fmt"

func main() {
	normalBuilder := GetBuilder("normal")
	iglooBuilder := GetBuilder("igloo")

	director := NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal house door type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal house window type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal house num of floors: %d\n", normalHouse.floors)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo house door type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo house window type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo house num of floors: %d\n", iglooHouse.floors)

	fmt.Println("\n\nExample 2")

	response := NewResponseBuilder().WithEmail("jane.doe@somemail.com").FromCity("Knowhere").WithBalance(1234.567).Build()
	fmt.Printf("Response: %+v\n", response)
}
