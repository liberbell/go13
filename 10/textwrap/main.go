package main

import (
	"fmt"
	"unicode"
)

func main() {
	const text = `チャットへようこそ！ご自身のプライバシーを守るとともに、YouTube のコミュニティ ガイドラインを遵守することを忘れないでください。Microsoft本社主催のMicrosoft Build 2021 が米国時間5月25日～27日にかけて開催されます。本イベントではBuildアップデートを含めた各サービスの最新アップデートをJAZUGの枠を超えて様々なコミュニティの方にお話ししていただきます！`

	const maxWidth = 40

	var lw int

	for _, r := range text {
		fmt.Printf("%c", r)

		// if lw++; lw > maxWidth {
		// 	// fmt.Printf(" <[%d]", lw)
		// 	lw = 0
		// 	fmt.Println()
		// } else if r == '\n' {
		// 	lw = 0
		// 	// fmt.Println()
		// }
		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}
}
