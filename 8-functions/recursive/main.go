package recursive

import "fmt"

func main() {
	answer := factorial(5)
	fmt.Println(answer)
}

func factorial(value int) int {
	if value == 1 {
		return value
	}

	// calling it's self. This func won't return until this inner on does
	return value * factorial(value-1)
}
