/**

もっとも簡単なオブジェクトの作成方法です。

C をつかっていた人にたいするわかりやすい説明としては、Object はメソッドをひもづけた構造体である。
ということです。

 */

package main

import "fmt"

type Dog struct {
    }
    func (self Dog) Berg() {
        fmt.Printf("Wan!\n");
    }

func main() {
    d := new(Dog)
    d.Berg()
}

