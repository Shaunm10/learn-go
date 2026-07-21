package main

func main() {

}

// numbers is the "rest of the parameters" much like spread in JavaScript
func sumUp(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
