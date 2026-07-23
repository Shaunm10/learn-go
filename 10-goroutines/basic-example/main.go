package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
}

// this func now takes a chan parameter to call
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	// tell's channel we are done.
	doneChan <- true
}

func main() {
	// declares a channel which is waiting for a boolean
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done) // pass the channel
	go greet("I hope you're liking the course!", done)

	// with 4 calls to the done channel, we will wait for 4 completions, so this code is more
	// concurrent

	// this will only work if the last task calls a 'done(chanChan)'
	// otherwise you will get an exception/error
	for doneChan := range done {
		fmt.Println(doneChan)
	}

	for range done {

	}

}
