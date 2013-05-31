package main

import (
    "fmt"
    "os"
    "strconv"
    "log"
)

func fib (n int) int {
    if n<2 {
        return 1
    } else {
        return fib(n-1) + fib(n-2)
    }
}

func main() {
    var n int;
    if len(os.Args) == 1 {
        n = 14
    } else {
        m, err := strconv.Atoi(os.Args[1])
        if err != nil {
            log.Fatal(err)
        }
        n = m
    }
    fmt.Printf("fib(%d)=%d\n", n,fib(n))
}

