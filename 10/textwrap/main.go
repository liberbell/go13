package main

import "fmt"

func main() {
	const text = `チャットへようこそ！ご自身のプライバシーを守るとともに、YouTube のコミュニティ ガイドラインを
遵守することを忘れないでください。Microsoft本社主催のMicrosoft Build 2021 が米国時間5月25日～27日にかけて開催
されます。本イベントではBuildアップデートを含めた各サービスの最新アップデートをJAZUGの枠を超えて様々なコミュニティ
の方にお話ししていただきます！`

	for _, r := range text {
		fmt.Printf("%c", r)
	}
}
