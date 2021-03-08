package main

import "fmt"

type Car1 struct {
	Name  string
	Price float32
}

type ElectricCar struct {
	Car1
	Battery int32
}

func main() {
	xp := ElectricCar{
		Car1{"xp", 200},
		70,
	}
	fmt.Println(xp.Name)

}
