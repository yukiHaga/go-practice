package main

import (
	"fmt"

	"github.com/yukiHage/go-practice/zoo/animals"
)

func main() {
    // 同じパッケージに定義した関数は、package_name.関数名()のようにしなくて良い
    fmt.Println(AppName())

    fmt.Println(animals.ElephantFeed())
    fmt.Println(animals.MonkeyFeed())
    fmt.Println(animals.RabbitFeed())
}