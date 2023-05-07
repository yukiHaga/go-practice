package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <-chan int) {
	  // 何も指定のない裸のforは無限ループ
    for {
				// チャネルからデータを取り出す
        i, ok := <-ch
				if !ok {
            // 受信できなくなったらループ終了
						// break文は最も近い位置のforループを中断させる
						fmt.Println(name + " is done.")
						break
				}
				fmt.Println(name, i)
		}
}

// mainが実行されたとき、mainでチャネルを生成して、ゴルーチンを３つ起動している
// main関数も一つのゴルーチンで実施されている
// go構文を用いて、任意の関数をゴルーチンとして起動することで、処理を並行して走らせている
func main() {
    // 明示的にバッファサイズを指定しない場合、バッファサイズ０のチャネルになる
    ch := make(chan int, 20)

    go receive("1st goroutine", ch)
    go receive("2nd goroutine", ch)
    go receive("3rd goroutine", ch)

    for i := 0; i < 100; i++ {
        ch <- i
    }
		close(ch)

		// ゴルーチンの完了を３秒待つ
		// timeパッケージは時間に関する処理を提供している
		// receive関数がゴルーチンとして起動している間も、main関数は先に進んでしまう。
		// そのため、main関数が終わらないように、receive関数側のゴルーチンの完了を待つために三秒停止している。
		// じゃないと、中途半端に数が出力されてプログラム全体が終了する
		// Go言語にはメインゴルーチンが終了したタイミングでプログラム全体を終了させるという特性がある
		time.Sleep(3 * time.Second)
}