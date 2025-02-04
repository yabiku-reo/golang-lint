package main

import "fmt"

func main() {
	fmt.Println("Hello CI!")

	// エラーハンドリングをしていないので、怒ってほしい
	hoge()
}

func hoge() error {
	return nil
}
