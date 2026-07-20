package main

import "fmt"

type Product struct {
	title string
	price float64
	id    int
}

func main() {
	// Time to practice what you learned!
	fmt.Println("ready to go....")

	// 1) Create a new array (!) that contains three hobbies you have
	hobbies := [3]string{"play guitar", "run", "play video games"}
	// 		Output (print) that array in the command line.
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Println(hobbies[0])
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	slice1 := hobbies[:2]
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice2 := hobbies[0:2]

	//fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice1 = slice1[1:3] // The GOTCHA is specifying index '3' which will make Go Look at the original array
	fmt.Println("updated slice1:", slice1)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"learn go", "build web Api's in go"}
	fmt.Println("Goals: ", courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Build a Rest Service in Go"
	courseGoals = append(courseGoals, "Add Go to resume")

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	productList := []Product{
		{
			title: "Fender Stratocaster",
			price: 999.99,
			id:    1,
		},
		{
			title: "Nintendo Switch 2",
			price: 499.99,
			id:    2,
		},
	}
	productList = append(productList, Product{
		title: "Weber Grill",
		price: 899.99,
		id:    2,
	})

	fmt.Println(productList)
	fmt.Println("done!")
}
