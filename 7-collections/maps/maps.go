package main

import "fmt"

func main() {

	fmt.Println("Maps testing ready to go...")

	websites := map[string]string{
		"Google": "http://google.com",
		"aws":    "https://aws.com",
	}
	fmt.Println(websites)

	// get the value from a key
	url := websites["aws"]

	// Adding a new element
	websites["cobalt"] = "www.cobalt.net"
}
