package main

import ( "fmt" )

// 普通の関数
func add(x int, y int) int {
    return x+y
}

// 多値をかえす関数
func add_sub(x int, y int) (int, int) {
    return x+y, x-y
}

func main() {
    // 関数よびだし
    fmt.Printf("%d\n", add(57, 32))

    // 多値をえる
    a, b := add_sub(57, 32)
    fmt.Printf("%d, %d\n", a, b)

    // こういうのはできないっぽい。なんかうまいことやる方法あるのかな?
    // fmt.Printf("%d, %d\n", add_sub(57, 32))
}

