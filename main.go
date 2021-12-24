package main

import (
	"fmt"

	"github.com/Nomankey/chancoin.git/person"
)


func main() {
	nomankey := person.Person{}
	nomankey.SeeDetails("chan", 25)
	fmt.Println("\nSee main",nomankey)
}