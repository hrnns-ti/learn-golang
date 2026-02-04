package main

import "fmt"

func main() {
	//NUMBERS
	//int	[int8, int16, int32, int64]					*include negative
	//uint	[uint8, uint16, uint32, uint64]				*exclude negative
	//float	[float32, float64, complex64, complex128]	*complex include imaginary numbers

	//alias	: byte = uint8
	//		: rune = int32
	//		: int  = int32 minimal
	//		: uint = uint32 minimal
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma lima = ", 3.5)

	//BOOLEAN
	fmt.Println("Boolean = ", true)
	fmt.Println("Boolean = ", false)
	fmt.Println("Boolean = ", 0 == 0)
	fmt.Println("Boolean = ", 0 != 0)

	//STRING
	var text string = "Hello World"
	var longtext = len(text)
	var alphabet = text[0:2]
	fmt.Println(text, longtext, alphabet)

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

	//DATA TYPE CONVERSION
	var num8 int8 = 127
	var num16 int16 = int16(num8)

	var num64 int32 = 123456789
	var num32 int16 = int16(num64) // can only convert into bigger type

	fmt.Println(num8, num16, num64, num32)

	//TYPE DECLARATION
	type num int
	var num1 num = 64
	fmt.Println(num1)
}
