package main

import (
	"fmt"
	// .を書くとパッケージ名の記述を省略できる
)

// 文字列かできることを示すインターフェース
// インターフェースを経由して型に定義してあるメソッドを利用するので、
// インターフェース自身はメソッドを呼び出すことで、どんな処理が実施されるのかを知らない(てか知らなくて良い)
// (メソッドの処理に関しては型の都合の処理も含まれるので、インターフェースは実装を知らなくて良い)
// そのため、実装した型ごとに、同じ名前だけど独自の処理を持つメソッドを呼び出せる
type Stringify interface {
    ToString() string
}

type Formatter interface {
    Format() string
}

type User struct {
    Name string
		Age int
}

// *Userという型に対してToStringメソッドを定義した
// つまり、*Userに対してインターフェースを実装したので、Userに対してインターフェースを実装したわけではない
func (u *User) ToString() string {
    return fmt.Sprintf("名前: %s, 年齢: %d", u.Name, u.Age)
}

func (u *User) Format() string {
    return fmt.Sprintf("名前: %s(%d)", u.Name, u.Age)
}

type Car struct {
    Name string
    Price int
}

func (c *Car) ToString() string {
    return fmt.Sprintf("名前: %s, 値段: %d", c.Name, c.Price)
}

func main() {
    user := User{Name: "yuki", Age: 26}
		car := Car{Name: "bmw", Price: 20000}

		// *Userに対してインターフェースを実装しているので、
		// *User型の値は、interface型(つまり、Formatter型)として扱える
		fmt.Println(format(&user)) // => 名前: yuki(26)

		a := [2]Stringify{&user, &car}

		for _, v := range a {
        fmt.Println(v.ToString())
		}
		/*
			名前: yuki, 年齢: 26
			名前: bmw, 値段: 20000
		*/
}

// f.ToString()を書くとエラーが起こる
// つまり、インターフェースを利用することで、利用しないメソッドへの依存を防げているということが分かる。
func format(f Formatter) string {
    return f.Format()
}