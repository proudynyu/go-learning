package main

import "fmt"

func intSeq() func() int {
    i := 0

    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())

    seq := func(a, b int) int {
        return a + b
    }
    
    a := seq(1, 2)

    fmt.Println(a)
}
