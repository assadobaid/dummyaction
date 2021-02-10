package calc

import "fmt"

func Add(a int, b int) int {
	fmt.Printf("Add is Called")
	return a + b
}

func Sub(a int, b int) int {
	fmt.Printf("Sub is Called, Hej, och hejdå")
	return a - b
}
