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

	tick := time.Tick(1*time.Second)

	for i:=10; i>0 ; i-- {
		fmt.Println(i)

		select {
		case <-abort:
			fmt.Println("Cancel launch")
			return
		case <-tick:
			
		}
	}
	

	launch3()
}

//!-


func launch3() {
	fmt.Println("Lift off!")
}