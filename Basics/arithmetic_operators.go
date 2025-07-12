package basics

import "fmt"

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a % b
	fmt.Println("Remainder", result)

}
