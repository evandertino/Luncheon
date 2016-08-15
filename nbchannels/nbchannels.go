package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non Blocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no messages received")
	}

	msg := "hi"
	// Non Blocking send
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Multi way Non Blocking select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
