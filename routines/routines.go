package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'j'; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\nDone")
}
