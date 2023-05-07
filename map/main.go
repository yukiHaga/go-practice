package main

import (
	"fmt"
)

func main() {
    m := make(map[string]string)
		m["name"] = "yuki"
		m["age"] = "26"
		fmt.Printf("%v\n", m) // => map[age:26 name:yuki]

		m2 := map[string]string{
		    "name": "yuki2",
        "age": "26",
		}
		fmt.Printf("%v\n", m2) // => map[age:26 name:yuki2]

		m3 := map[string]string{
        "name": "yuki",
		}
		fmt.Printf("%v\n", m3) // => map[name:]

	  v1, ok1 := m3["name"]
		v2, ok2 := m3["age"]
		fmt.Printf("%v, %v\n", v1, ok1) // => yuki, true
		fmt.Printf("%v, %v\n", v2, ok2) // => "", value

}