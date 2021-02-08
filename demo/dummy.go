package main

import (
	"fmt"

	calculator "github.com/assadobaid/dummyaction"
)

func main() {
	fmt.Println("Add(6, 3) =>", calculator.Add(6, 3))
	fmt.Println("Subtract(6, 3) =>", calculator.Subtract(6, 3))
}
