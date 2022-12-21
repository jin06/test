package main

import "fmt"

type Animal interface {
	Eat()
}
type Dog struct {
}

func (d *Dog) Eat() {
	fmt.Println("Dog eat")
}
func main() {
	var d Animal = &Dog{}
	val, ok := d.(*Dog)
	fmt.Println(ok)
	val.Eat()
}
