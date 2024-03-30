package main

import (
	"fmt"
	"math/rand"
	"time"
)

func async_fn(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i);
		var amt = time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// go routines are lightweight way to do concurrency, think of it as close to async/await, but there is no coloring of functions, and its closer to doing Threads.

// To execute a function as a go routine - on another thread, execute the function with the keyword "go" before the function execution

func main() {
	for i := 0; i < 10; i++ {
		go async_fn(i)
	}
	var input string
	fmt.Scanln(&input)
}