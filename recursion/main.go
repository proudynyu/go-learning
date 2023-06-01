package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }

    return n * fact(n - 1)
}

func fib(n int) int {
    if n < 2 {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}

func fizz_buzz(n int) {
    for i := 0; i <= n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("fizzbuzz")
            continue
        } else if i % 3 == 0 {
            fmt.Println("fizz")
        } else if i % 5 == 0 {
            fmt.Println("buzz")
        } else {
            fmt.Println(i)
        }
    }
}

func main() {
    fmt.Println(fact(7))

    fmt.Println(fib(7))
    fizz_buzz(10)
}
