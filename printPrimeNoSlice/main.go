/*
 With the help of go routine we have to prine the prime nos as per use input
 1 => 1
 5 => 5 prime nos
*/

package main

import (
	"fmt"
)

func isPrime(no int) bool {
	if no < 2 {
		return false
	}

	for i := 2; i*i <= no; i++ { // i*i <= no is faster than i <= no/2 or i <= no as it reduces the number of iterations
		if no%i == 0 {
			return false
		}
	}

	return true
}

func findPrimeNos(input int, ch chan int) {
	index := 2
	found := 0

	for {
		if isPrime(index) { // check if no is prime
			ch <- index // send the prime no to channel
			found++     // increment the found prime no
		}

		if found == input { // check if we have found the required prime nos
			break // if yes then break the loop
		}

		index++ // increment the index to check next no
	}

	close(ch)
}

func main() {
	var input int
	fmt.Println("Enter the number: ")
	fmt.Scan(&input)
	ch := make(chan int)

	go findPrimeNos(input, ch)

	for no := range ch {
		fmt.Println(no)
	}
}
