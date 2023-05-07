package main

import (
	"fmt"
	"math"
)

func main() {
	  i := 1
		fmt.Printf("数字: %d\n", i)

		i = 2
		fmt.Printf("数字: %d\n", i)

		m := 256
		b := byte(m)
		fmt.Printf("数字: %d\n", b)

		fmt.Println(doSomething(uint32(4), uint32(3)))

		a := [4]int{1, 2, 3, 4}
		fmt.Printf("配列: %v\n", a) // => 配列: [1 2 3 4]

		var (
			a1 = [3]int{4, 5, 6}
			a2 = [3]int{1, 2, 3}
		)

		a2 = a1
		fmt.Printf("配列a2: %v\n", a2)

		a2[0] = 0
		a2[2] = 0

		fmt.Printf("配列a2: %v\n", a2)
		fmt.Printf("配列a1: %v\n", a1)

    hello()

		f := func(a, b int) int { return a + b }

		fmt.Printf("値: %d\n", f(2, 5)) // => 値: 7
		fmt.Printf("型: %T\n", f) // => 型: func(int, int) int
}

// Goでは戻り値の型は必須。TSでは省略可能
// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できる。TSの場合、全ての引数に型を記述しないといけない
// uint32型が表現できる整数の最大値は４２億弱
func doSomething(a, b uint32) bool {
	// 普通にa + bの不等式を作ると、オーバーフローしてラップアラウンドしてしまうので、引き算の不等式にしている。
	// 蓬莱ならa + bで最大値より大きな数になって惜しいのだが、そのまま不等式にすると、オーバーフローしてラップアラウンドしてしまう
	if (math.MaxInt32 - a) < b {
		return false
	} else {
		return true
	}
}

func hello() {
    fmt.Println("Hello")
}