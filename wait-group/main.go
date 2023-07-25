package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go div(100, i, &wg)
	}
	wg.Wait()
}

func div(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := a / a // Error here
	fmt.Printf("%d/%d=%d\n", a, b, res)
}
