package main

import "fmt"

func main(){
	foods := []string{"food", "pizza", "pasta", "ramen"}
	fmt.Printf("%v\n", foods)
	newFoods := append(foods, "tomato")
	fmt.Printf("%v\n", newFoods)
}