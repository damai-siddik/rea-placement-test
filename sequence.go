package main

import "fmt"

func Sequence(n int) []int {
    // Store the sequence here
    numbers := make([]int, n)

    // Start values
    numbers[0] = 3
    numbers[1] = 5

    for i := 2; i < n; i++ {
        numbers[i] = numbers[i-1] + i * 2
    }

    return numbers
}

func main() {
    fmt.Println(Sequence(100));
    
}
