package main

import "fmt"

func main() {
	const text = `チャットへようこそ！ご自身のプライバシーを守るとともに、YouTube のコミュニティ ガイドラインを遵守することを忘れないでください。`

	for _, r := range text {
		fmt.Printf("%c", r)
	}
}
