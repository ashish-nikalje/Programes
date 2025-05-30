package main

import (
	"fmt"
	"sync"
)

func printEven(input uint, evenChan, oddChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= int(input); i += 2 {
		<-evenChan
		fmt.Println("Even: ", i)
		oddChan <- true
	}
}

func printOdd(input uint, evenChan, oddChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= int(input); i += 2 {
		<-oddChan
		fmt.Println("Odd: ", i)
		evenChan <- true
	}
}

func main() {
	fmt.Print("Hello, Print How nos you want to print in Sequnce?: ")
	var input uint
	fmt.Scanln(&input)

	evenChan, oddChan := make(chan bool, 1), make(chan bool, 1) // Buffered channel to avoid deadlock
	defer close(evenChan)
	defer close(oddChan)

	var wg sync.WaitGroup

	wg.Add(2)
	go printEven(input, evenChan, oddChan, &wg) // Start the sequence
	go printOdd(input, evenChan, oddChan, &wg)  // Start the sequence

	oddChan <- true // Start the sequence with odd number first

	wg.Wait()
}
