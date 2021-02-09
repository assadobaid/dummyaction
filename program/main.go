package main

import (
	"fmt"

	calc "github.com/assadobaid/dummyaction"
)

func main() {
	fmt.Println("Add(6, 3) =>", calc.Add(6, 3))
	fmt.Println("Subtract(6, 3) =>", calc.Sub(6, 3))
}
