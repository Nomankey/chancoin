package person

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) SeeDetails (name string, age int) {
	p.name = name
	p.age = age
	fmt.Printf("See detaills, name %s, age %d", p.name, p.age)
}
