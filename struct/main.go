package main

import (
	"fmt"
	// .を書くとパッケージ名の記述を省略できる
	. "github.com/yukiHaga/go-practice/struct/user"
)

func main() {
    type MyInt int
		var age MyInt = 26
		fmt.Println(age) // => 26

		var u User
		fmt.Println(u.FirstName) // => ""
		fmt.Println(u.Age) // => 0

		u.FirstName = "yuki"
		u.Age = 26

		fmt.Println(u.FirstName) // => "yuki"
		fmt.Println(u.Age) // => 26

		// 定義した構造体型を利用して、構造体を生成している
		r := User{FirstName: "yuki", Age: 26}
		fmt.Println(r) // => {yuki 26}

		u2 := User{FirstName: "yuki", LastName: "hoge", Age: 26}
		fmt.Println(u2.FullName()) // => yuki haga
		fmt.Println(u2.Age) // => 26

		(&u2).UpdateAge(27)
		fmt.Println(u2.Age) //=> 2

		fmt.Println(NewUser("yuki", "hoge", 26, false))

		fmt.Println(u2)
}