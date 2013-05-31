/**
 * foreach 的な。
 */

package main

import (
    "fmt"
)

func main() {
    // 配列を代入する
    ary := [...]int{5963,4649};

    // 配列の中身を順番にみる
    for i, e := range ary {
        fmt.Printf(" %d, %d\n", i, e)
    }

    // 添字がいらないときは _ をつかう
    for _, e := range ary {
        fmt.Printf(" %d\n", e)
    }

    // 配列サイズを表示する
    fmt.Printf(" Size: %d\n", len(ary))
}

