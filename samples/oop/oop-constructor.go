/*

コンストラクタの作り方です。コンストラクタは普通の関数として定義します。

*/

package main

import "fmt"

type Person struct {
    name string
}
func NewPerson(name string) *Person {
    p := new(Person)
    p.name = name
    return p
}
func (self Person) greeting() {
    fmt.Printf("Hi! I'm %v\n", self.name)
}

func main() {
    p := NewPerson("John")
    p.greeting()
}

