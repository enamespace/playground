package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) PointerPrint() {
	fmt.Println("this is point receiver for " + p.name)
}

func (p Person) ValuePrint() {
	fmt.Println("this is value receiver for " + p.name)
}

func main() {
	p := &Person{name: "pointer"}
	v := Person{name: "value"}

	p.PointerPrint()
	(*p).PointerPrint()
	v.ValuePrint()
	(&v).ValuePrint()
	(Person{name: "value"}).ValuePrint()
}
