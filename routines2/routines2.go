package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'j'; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait()

	fmt.Println("\nDone")
}
