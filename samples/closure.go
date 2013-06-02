/**
 * クロージャーの例です。クロージャーとは、変数をバインドしてもつ無名関数のことです。
 * Perl, Ruby などの LL とおなじ感覚でクロージャーを利用することができます。
 */
package main

import (
    "fmt"
)

func counter() (func () int) {
    i := 0
    return func () int {
        i = i+1
        return i
    }
}

func main() {
    c := counter()
    fmt.Printf("%d\n", c());
    fmt.Printf("%d\n", c());
    fmt.Printf("%d\n", c());
    fmt.Printf("%d\n", c());
}

