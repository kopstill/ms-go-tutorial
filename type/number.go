package main

import (
	"fmt"
	"math"
)

var integer8 int = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807

func main() {
	fmt.Println(integer8, integer16, integer32, integer64)
	// fmt.Println(integer16 + integer32)

	rune := 'G'
	fmt.Println(rune)

	// var integer32 int32 = 2147483648
	// fmt.Println(integer32)

	var integer32 int = 2147483648
	fmt.Println(integer32)

	// var integer uint = -10
	// fmt.Println(integer)

	var integer int = 10
	fmt.Println(integer)

	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	fmt.Println(float32, float64)

	fmt.Print(math.MaxFloat32, math.MaxFloat64, "\n")

	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34
	fmt.Println(e, Avogadro, Planck)
}
