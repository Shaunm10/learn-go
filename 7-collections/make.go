package main

import "fmt"

func main() {

	makeWithArray()
	makeWithMap()

}

func makeWithArray() {
	// only has 2 "free" spaces, but will need to create a new array
	// if more items are added
	namesJustWithTwo := [2]string{}

	// this array will be created with 2 "free" spaces, but allocated with 5 so Go
	// doesn't have to recreate an array when expanded
	namesWithMoreSpace := make([]string, 2, 5)
	namesJustWithTwo[0] = "Maggie"
	namesWithMoreSpace[0] = "Jack"

	// 2
	fmt.Println(namesJustWithTwo, cap(namesJustWithTwo))

	//5
	fmt.Println(namesWithMoreSpace, cap(namesWithMoreSpace))

	// loop through the items in an array
	for index, value := range namesJustWithTwo {
		// executes on each item
		fmt.Println(value)
		fmt.Println(index)
	}
}

func makeWithMap() {

	type floatMap map[string]float64
	courseRatings := make(floatMap, 6)

	fmt.Println(courseRatings)

	for _, value := range courseRatings {
		// iterate through each item in the map
		fmt.Println(value)
	}
}
