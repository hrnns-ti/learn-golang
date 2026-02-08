package main

import "fmt"

func main() {
	//VARIABLE
	var name1 string
	name1 = "Jack"
	var age1 int = 18

	name2 := "Jack"
	age2 := 18

	var (
		name3 string = "Jack"
		age3         = 18
	)

	fmt.Println(name1, age1, name2, age2, name3, age3)

	//CONSTANT
	const month1 string = "January"
	const day1 int8 = 1

	month2 := "February"
	day2 := 2

	const (
		month3 string = "March"
		day3          = 3
	)

	fmt.Println(month1, day1, month2, day2, month3, day3)
}
