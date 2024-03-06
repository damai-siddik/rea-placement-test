package main

import (
	"fmt"
	"math"
)

func Sequence(n int) []int {
    // Store the sequence
    numbers := make([]int, n)

    // Start values
    numbers[0] = 3
    numbers[1] = 5

    for i := 2; i < n; i++ {
        numbers[i] = numbers[i-1] + i * 2
    }

    return numbers
}

func RecursiveRandomNumberGenerator(n int, seed int) int {
     // Define constants
     a := 1103
     c := 1234
     m := int(math.Pow(2, 31))
 
     // Check if n is 0 (base case)
    if n == 0 {
        return seed
    }

    // Make recursive call
    previousNumber := RecursiveRandomNumberGenerator(n-1, seed)
 
    // Calculate new random number
    newNumber := (a * previousNumber + c) % m

    return newNumber
}

func main() {
    fmt.Println(Sequence(100));
    fmt.Println(RecursiveRandomNumberGenerator(100, 1));
}
