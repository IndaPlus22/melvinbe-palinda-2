// Doesn't work:
/*
package main

import (
	"fmt"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}
*/

// Why not:
/*
The main goroutine sends 11 integers to the channel and then closes it, without waiting for the Print
goroutine to finish processing the channel. The program then exits before the integers are printed from the channel.
*/

// Works:
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	// Wait for 1 goroutine
	wg.Add(1)

	go Print(ch, wg)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	// Wait for integers to be printed
	wg.Wait()
}

func Print(ch <-chan int, wg *sync.WaitGroup) {
	for n := range ch {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(n)
	}
	// Ok now the program can exit
	wg.Done()
}

// Why:
/*
Using a waitgroup, the main goroutine can wait until the Print goroutine is done printing the integers before exiting
*/
