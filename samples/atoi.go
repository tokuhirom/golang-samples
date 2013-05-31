package main

import (
    "fmt"
    "strconv"
    "log"
)

func main() {
    i, err := strconv.Atoi("5963")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d\n", i)
}

