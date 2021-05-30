package main

import (
	"fmt"
	"unicode"
)

func main() {
	// const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
	// Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
	// Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`
	const text = `チャットへようこそ！ご自身のプライバシーを守るとともに、YouTube のコミュニティ ガイドラインを遵守することを忘れないでください。Microsoft本社主催のMicrosoft Build 2021 が米国時間5月25日～27日にかけて開催されます。本イベントではBuildアップデートを含めた各サービスの最新アップデートをJAZUGの枠を超えて様々なコミュニティの方にお話ししていただきます！`

	const maxWidth = 40

	var lw int // line width

	for _, r := range text {
		fmt.Printf("%c", r)

		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}
	fmt.Println()
}
