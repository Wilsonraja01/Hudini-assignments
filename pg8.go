// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.
 
// Task:
 
// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.
package main
import(
	"fmt"
	// "sync"
)

	// -------------------------------------------------------------------------------------
	 
	// Objective:
	// Learn how to use the select statement to handle multiple channels.
	 
	// Task:
	 
	// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
	// Use a select statement in the main function to receive messages from both channels and print them.
	func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	 
	channel1Closed := false
	channel2Closed := false
	 
	go func() {
	channel1 <- "sending message from channel 1\n"
	close(channel1)
	}()
	 
	go func() {
	channel2 <- "sending message from channel 2\n"
	close(channel2)
	}()
	for {
	select {
	case msg1, ok := <-channel1:
	if !ok {
	channel1Closed = true;
	} else {
	fmt.Println(msg1)
	}
	case msg2, ok := <-channel2:
	if !ok {
	channel2Closed = true;
	} else {
	fmt.Println(msg2)
	}
	}
	 
	if channel1Closed && channel2Closed {
	break;
	}
	}
	}