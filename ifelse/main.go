package main

import "fmt"

func is_even(num int) bool {
    return num % 2 == 0
}

func main() {
    for n := 0; n <= 10; n++ {
        if is_even(n) {
            fmt.Printf("%d is even\n", n) 
        } else {
            fmt.Printf("%d is odd\n", n) 
        }
    }
}
