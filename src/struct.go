package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%d %s", p.Age, p.Name)
}

func (p *Person) Hi() {
	fmt.Println("My Name Is" + p.Name)
}

func main() {
	p := &Person{"Pompeu", 33}
	p.Hi()
	fmt.Println(p)
}
