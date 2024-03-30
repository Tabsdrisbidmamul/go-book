package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0 ;  ; i++ {
		c <- "ping"
	}
}

func printer(c <-chan string) {
	for {
		// msg := <- c
		fmt.Println(<- c)
		time.Sleep(time.Second * 1)
	}
}

func ponger (c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// Channels
// channels are used to share information across go routines. The keyword chan followed by the type
// "chan string", meaning that the channel will receive a string

// the syntax <- (left arrow operator) is used to send and receive messages on the channel, so 
// c <- "ping" means send "ping"
// msg := <- c means receive the message and store it in msg

// Synchronisation
// using a channel synchronises the go routines, so when pinger attempts to send a message on the channel, it will be blocked until printer is ready to receive it. Amazing that we don't need to create a lock for the mutex/ semaphore anymore

// Channel Direction
// we can restrict the channel direction by changing the function signature

// If we only want the function to send we write it as "c chan<- string" the spacing on the left arrow operator is important, there should be no space between the keyword "chan" and the left arrow operator

// If we want the function to receive the message, we write "c <-chan string", means that the function can only receive a message

// If we don't specify the direction, then the channel is bi-directional


func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}