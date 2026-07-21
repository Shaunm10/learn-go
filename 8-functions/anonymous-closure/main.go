package anonymousclosure

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// anonymous function example
	transformed := transformNumbers(&numbers, func(numb int) int {
		return numb * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {

	return func(num int) int {
		// because we can use the factor parameter is due to Closures
		// it's in scope.
		return num * factor
	}
}
