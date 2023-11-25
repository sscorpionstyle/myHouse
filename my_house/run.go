package main

import (
	"fmt"
)

func showHouseInfo(house House) {
	fmt.Println("House Address:", house.Address)
	fmt.Println("House Square:", house.Size.Width*house.Size.Length, "meters")
	fmt.Println("Family Members:")
	for _, member := range house.Family {
		fmt.Println(member.Role, member.Name, ",", member.Age, "y.o.")
	}
	fmt.Println("\nFurniture:")
	for _, furniture := range house.Furniture {
		fmt.Println(furniture.Name, " - ", furniture.Color, ",")
	}
	fmt.Println("\nAppliances:")
	for _, appliance := range house.Appliances {
		fmt.Println(appliance.Name, "in the", appliance.Room)
	}
}

func runProject() {
	house := NewHouse("Yukhnov, Uritskogo street, 104", 10.5, 15.2)
	familyMembers := []FamilyMember{
		{"Tanya", 51, "Mom"},
		{"Andrey", 52, "Dad"},
		{"Anzhelika", 20, "Daughter"},
		{"Archie", 9, "Dog"},
	}
	furniture := []Furniture{
		{"Sofa", "Brown"},
		{"Bed", "White"},
		{"Dining Table", "Yellow"},
		{"Wardrobe", "White"},
		{"Billiard", "Green"},
	}
	appliances := []Appliance{
		{"Fridge", "Kitchen"},
		{"Microwave", "Kitchen"},
		{"Kettle", "Kitchen"},
		{"TV", "Living Room"},
		{"Laptop", "Any Room"},
	}

	house.Family = familyMembers
	house.Furniture = furniture
	house.Appliances = appliances

	showHouseInfo(house)
}
