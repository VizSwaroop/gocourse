package basics

import "fmt"

func main() {

	//Simple iteration over Range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)

	}

	// iterate over a collection

	numbers := []int{1, 2, 3, 4, 5, 6}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value : %d\n", index, value)
	}

}
