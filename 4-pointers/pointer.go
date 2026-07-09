package main

import "fmt"

func main() {

	age := 32

	var agePointer *int

	// note the referencing "&"
	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	adultAge := getAdultYears(age)

	//14
	fmt.Println("The adult age is now:", adultAge)

	updateAgeToAdultYears(agePointer)

	// 14
	fmt.Println("The  age is now:", age)

}

// pass by value
func getAdultYears(age int) int {
	return age - 18
}

// pass by reference
func updateAgeToAdultYears(age *int) {
	*age = *age - 18
}
