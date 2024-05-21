// Objective:
// Learn how to use a buffered channel.

// Task:

// Write a program that creates a buffered channel with a capacity of 3.
// The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// Then, receive and print the integers from the channel.

package main

import "fmt"
func main(){
	msg:=make(chan int,3)
	msg<- 1
	msg<- 2
	msg<- 3
	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}