package main

import (
	"fmt"
	"math"
)

const str string = "Constantes"

func main() {
	fmt.Println(str)

	const max = 500000000

	const dig = 3e20 / max
	fmt.Println(dig)

	fmt.Println(int64(dig))

	fmt.Println(math.Sin(max))
}
