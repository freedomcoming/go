package main

import "fmt"

type Car struct {
	Name  string
	Price float32
}

func (car *Car) Info() {
	fmt.Printf("%v price: [%v]", car.Name, car.Price)
}

func main() {
	car := Car{
		Name:  "BMW",
		Price: 100.0,
	}
	car.Info()

}
