package otherpackage

import "fmt"

type Person struct {
	Name string
	Age  int
	Fre  []Friend
}
type Friend struct {
	Anime string
	Aage  int
}

func (p *Person) PersonInfo(action string) (res string) {
	res = "ok"
	fmt.Printf("%v,%v ,%v", p.Name, p.Age, action)
	return

}
