package main

import (
    "fmt"
    "math"
)

func main() {
	numbers := []int{16,8,42,4,56,9}

    max := math.MinInt16

    for i := range numbers {
        if numbers[i] > max {
            max = numbers[i]
        }
    }
    fmt.Println(max)
}
