// Objective:
// Understand how to use channels for communication between goroutines.

// Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.
package main

import (
	"fmt"
	"sync"
	"time"
)

 func main(){
	msg:=make(chan int)
	wg:=new(sync.WaitGroup)
	go sendNum(msg,wg)
	for i:=1;i<=10;i++{
		fmt.Println(<-msg)
	}
	wg.Wait()
 }

 func sendNum(ch chan int,wg1 *sync.WaitGroup) {
	defer wg1.Done()
	for i:=1;i<=10;i++{
		ch<-i
		time.Sleep(time.Second)
	}
 }