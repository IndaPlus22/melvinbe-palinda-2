// Doesn't work:
/*
package main

import "fmt"

func main() {
	ch := make(chan string)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
*/

// Why not:
/*
The program is currently causing a deadlock. When "hello world!" is sent to the channel, the main goroutine is
blocked until another goroutine recieves from the channel. But since there is only one goroutine, it will never unblock.
*/

// Works:
package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() { ch <- "Hello world!" }()
	fmt.Println(<-ch)
}

// Why:
/*
If "hello world!" instead is sent from  another goroutine using the go keyword on a new function, the main goroutine
will not be deadlocked. It can then recieve from the channel and print on the next line.
*/
