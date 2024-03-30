package main

import (
	"fmt"
	"time"
)

// select - like a switch but for channels
// the select statement will "select" a value from the channel when its ready to be received (or sends from it)

// In this program, c1 will send a message every 2 seconds and c2 will send a message every 3 seconds
// The select will choose a cases
// c1 is ready, will choose c1
// c2 is ready, will choose c2
// when both are ready, will choose either of them at random
// if none are ready, then it will block, until one of the channels becomes available

// Buffered Channels
// channels are synchronous by default, so it will block and wait to send and receive
// By specifying a buffer, it will make the channel asynchronous, so it will no longer block to send or receive
// To do so, you specify another parameter when making a channel

// var c = make(chan int, 1)
// The channel above has a capacity of 1, when its full the channel will be blocked. However since there is a buffer now, if there were 2 or more go routines attempting to access the above channel, there will be a deadlock

func main() {
	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
			// default:
			// 	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}