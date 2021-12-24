package main

import "fmt"

type person struct {
	name string
	age int
}

func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s, and I'm %d years old", p.name, p.age)
}

func main() {
	nomankey := person{name: "chan", age: 17}
	nomankey.sayHello()
}