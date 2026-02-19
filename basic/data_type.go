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

	//ARRAY
	var name [2]string
	name[0] = "Alice"
	name[1] = "Bob"
	fmt.Println(name[0], name[1])
	fmt.Println(name)

	var values = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(values)

	longName := len(name) //LENGTH AN ARRAY
	nValues := values[0]  // DATA IN ARRAY INDEX N
	values[1] = 100       // CHANGE A DATA IN INDEX N

	fmt.Println(longName)
	fmt.Println(nValues)
	fmt.Println(values)

	//SLICE
	moon := [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}
	slice := moon[2:5]
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice)
	fmt.Println(moon)
}
