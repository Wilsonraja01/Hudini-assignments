// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:

// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.

package main

import "fmt"



func main(){
	channel1:=make(chan string)
	channel2:=make(chan string)

	go func ()  {
		channel1<-"first message"
	}()
	go func ()  {
		channel2<-"second message"
	}()
	for{
		select {
		case a:=<-channel1:
			fmt.Println(a)
		case b:=<-channel2:
			fmt.Println(b)
			return
		}
	}

}