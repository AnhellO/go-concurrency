package main

import (
	"fmt"
	"time"
)

func doSomething() <-chan struct{} {
	done := make(chan struct{})

	go func() {
		timer := time.Tick(1 * time.Second)

		i := 0
		for {
			select {
			case <-timer:
				fmt.Println("tick!")
				i++
				if i >= 5 {
					done <- struct{}{}
				}
			}
		}
	}()

	return done
}

func main() {
	done := doSomething()
	<-done
}
