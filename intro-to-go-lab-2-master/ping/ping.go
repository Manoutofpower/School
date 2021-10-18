package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
	for {
		msg := "ping"
		fmt.Printf("Foo is sending :%s \n", msg)
		channel <- msg
		fmt.Printf("Foo has  received :%s \n", <-channel)
	}
}

func bar(channel chan string) {
	// TODO: Write an infinite loop of receiving "pings" and sending "pongs"
	for {
		msg := <-channel
		fmt.Printf("Bar has received:%s \n", msg)
		msg = "pong"
		fmt.Printf("Bar is sending :%s \n", msg)
		channel <- msg
	}
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	ch := make(chan string)
	go foo(ch) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(ch)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()

	select {
	case <-time.After(1 * time.Second):
		return
	}
}
