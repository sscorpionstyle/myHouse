package main

type House struct {
	Address    string
	Size       HouseSize
	Family     []FamilyMember
	Furniture  []Furniture
	Appliances []Appliance
}

type HouseSize struct {
	Width  float64
	Length float64
}

func NewHouse(address string, width, length float64) House {
	size := HouseSize{Width: width, Length: length}
	return House{Address: address, Size: size}
}
