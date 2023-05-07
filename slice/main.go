package main

import (
	"fmt"
)

func main() {
    s := []int{1, 2, 3, 4}
		fmt.Printf("%v\n", len(s)) // => 4

		s = append(s, 5)
		fmt.Printf("%v\n", s) // => [1 2 3 4 5 7 6]

		s = append(s, 7, 6)
		fmt.Printf("%v\n", s) // => [1 2 3 4 5]

		s2 := []int{2, 4, 6}
		s3 := []int{3, 6, 9}

		s4 := append(s2, s3...)
		fmt.Printf("%x\n", s4) // => [2 4 6 3 6 9]

		fmt.Printf("%d\n", sum(2, 3, 5)) // => 10

		var (
        a [3]int
        s5 []int
		)

		fmt.Printf("%v\n", a)
		fmt.Printf("%v\n", s5)
}

func sum(a ...int) int {
    s := 0

    for _, v := range a {
        s += v
		}

		return s
}
