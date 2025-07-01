package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// syscall.SIGINT should be cmd + c
	// syscall.SIGTERM should be kill this process
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("received signal: ", sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("done")
}
