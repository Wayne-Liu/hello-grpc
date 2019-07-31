package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown, press enter to abort")

	select {
	case <-time.After(10*time.Second):

	case <-abort:
		fmt.Println("Launch aborted!")
		return
		
	}

	launch2()
}

//!-


func launch2() {
	fmt.Println("Lift off!")
}