package main

func main() {

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {

	returnValues := []int{}

	for _, val := range *numbers {
		returnValues = append(returnValues, transform(val))
	}

	return returnValues
}

func doubler(num int) int {
	return num * 2
}
func getTransformerFunction() func(int) int {

	return
}
