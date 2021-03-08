package main

import "fmt"

type Car2 struct {
	Name  string
	Price float32
}
type ElectricCar1 struct {
	Car2
	Battery int32
}

type PetrolCar struct {
	Car2
	Gasoline int32
}

//定义一个接口
type RunService interface {
	Run()
}

// 实现1
func (car PetrolCar) Run() {
	fmt.Printf("%s PetrolCar run \n", car.Name)
}

// 实现2
func (car ElectricCar1) Run() {
	fmt.Printf("%s ElectricCar1 run \n", car.Name)
}

func Do(run RunService) {
	run.Run()
}

func main() {
	xp := ElectricCar1{
		Car2{Name: "xp", Price: 200},
		70,
	}
	petrolCar := PetrolCar{
		Car2{Name: "BMW", Price: 300},
		50,
	}
	Do(xp)
	Do(petrolCar)

}

//xp ElectricCar1 run
//BMW PetrolCar run
