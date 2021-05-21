package article

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

// ParseArticle procides Article's behavior
type ParseArticle struct{}

// OGPScanResult is OGP Meta Tag value
type OGPScanResult struct{
	Title       	string
	Description 	string
	ThumbnailURL 	string
	Header      	string
	Body        	string
}

// ScanArticle is scan article
func ScanArticle(url string) OGPScanResult{
	
	fmt.Println(url)
	// URLの静的解析を実行
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}


	var pageTitle string
	var metaDescription string
	var metaImageURL string

	pageTitle = doc.Find("title").Contents().Text()
	// 実装参考 https://qiita.com/Yaruki00/items/b50e346551690b158a79
	doc.Find("meta").Each(func(index int, item *goquery.Selection) {

		// Attr("rtn true", "rtn false")
			if( item.AttrOr("property","") == "og:description") {
					metaDescription = item.AttrOr("content", "")
			}

			if( item.AttrOr("property","") == "og:image") {
				metaImageURL = item.AttrOr("content", "")
		}
	})
	fmt.Printf("Page Title: '%s'\n", pageTitle)
	fmt.Printf("Meta Description: '%s'\n", metaDescription)
	fmt.Printf("Meta Image: '%s'\n", metaImageURL)

	body := doc.Find("body")
	fmt.Println(body)

	// OGP解析で取得したメタ情報
	rtnVal := OGPScanResult{
		Title					: pageTitle,
		Description		:	metaDescription,
		ThumbnailURL	: metaImageURL,
		Header				: "ほげほげ",
		Body					: "ふがふが",
	}


	return rtnVal

}