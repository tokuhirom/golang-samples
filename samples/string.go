package main

import (
    "fmt"
//  "os"
//  "io/ioutil"
//  "log"
    "strings"
)

func main() {
    // 文字列の連結
    fmt.Printf("%s\n", "x" + "y")

    // 文字列の一部をとりだす
    fmt.Printf("%c\n", "abcdefg"[0])
    fmt.Printf("%s\n", "abcdefg"[0:3])
    fmt.Printf("%s\n", "abcdefg"[3:])
    fmt.Printf("%s\n", "abcdefg"[:4])

    // 大文字にする
    fmt.Printf("%s\n", strings.ToUpper("aBc"))

    // 小文字にする
    fmt.Printf("%s\n", strings.ToLower("aBc"))

    // 連結する
    ary := []string{"foo","bar","baz"}
    fmt.Printf("%s\n", strings.Join(ary, ","))
}

