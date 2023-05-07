package main

import (
	"fmt"
)

func updateNum(a [3]int) {
    a[0] = 7
    a[1] = 7
    a[2] = 7
}

func updateNumOfReference(p *[3]int) {
    (*p)[0] = 7
    (*p)[1] = 7
    (*p)[2] = 7
}

func main() {
    var i int
		// ポインタpを定義
		p := &i
		i = 5
		fmt.Println(*p) // => 5
		*p = 10
		fmt.Println(i) // => 10

		a := [3]int{1, 2, 3}
		updateNum(a)
		fmt.Printf("配列a: %d\n", a)

		b := &a
		fmt.Printf("配列b: %d\n", *b) // => 配列b: [1 2 3]
		updateNumOfReference(b)
		fmt.Printf("配列b: %d\n", *b) // => 配列b: [7 7 7]
}

