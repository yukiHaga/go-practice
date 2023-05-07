package main

import (
	"fmt"
)

func main() {
	  a := 1
    switch a {
		case 1:
			fmt.Printf("%d\n", a)
		default:
			fmt.Println("no hit")
		}

		b := 2
    switch {
		case a == 1 && b == 2 :
			fmt.Printf("%d\n", a + b)
		default:
			fmt.Println("no hit")
		}

		// tsだと変数名の後ろに;が必要
		var i interface{} = 3

		d, isInt := i.(int)
		e, isFloat65 := i.(float64)
		f, isString := i.(string)

		fmt.Printf("%d, %v\n", d, isInt) // => 3, true
		fmt.Printf("%v, %v\n", e, isFloat65) // => 0, false
		fmt.Printf("%s, %v\n", f, isString) // => "", false

		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("dont know")
		}
}