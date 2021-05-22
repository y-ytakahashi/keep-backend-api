package services

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParseArticle procides Article's behavior
type ParseArticle struct{}

// OGPScanResult is OGP Meta Tag value
type OGPScanResult struct{
	URL				string
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
	var metaURL string

	pageTitle = doc.Find("title").Contents().Text()

	// 以下の指定で記事本文の取得が可能
	pageHeader := doc.Find("header").Contents().Text()
	cnvPageHeader := convNewline(pageHeader, " ")

	pageBody := doc.Find("body").Contents().Text()
	cnvPageBody := convNewline(pageBody," ")


	// 実装参考 https://qiita.com/Yaruki00/items/b50e346551690b158a79
	doc.Find("meta").Each(func(index int, item *goquery.Selection) {

		// Attr("rtn true", "rtn false")
		if( item.AttrOr("property","") == "og:description") {
				metaDescription = item.AttrOr("content", "")
		}

		if( item.AttrOr("property","") == "og:image") {
			metaImageURL = item.AttrOr("content", "")
		}

		if( item.AttrOr("property","") == "og:url") {
			metaURL = item.AttrOr("content", "")

			if metaURL == "" {
				metaURL = url
			}

		}
	})

	fmt.Printf("Page Title: '%s'\n", pageTitle)
	fmt.Printf("Meta Description: '%s'\n", metaDescription)
	fmt.Printf("Meta Image: '%s'\n", metaImageURL)


	// OGP解析で取得したメタ情報
	rtnVal := OGPScanResult{
		URL						: metaURL,
		Title					: pageTitle,
		Description		:	metaDescription,
		ThumbnailURL	: metaImageURL,
		Header				: cnvPageHeader,
		Body					: cnvPageBody,
	}



	return rtnVal

}
// convNewline 全ての改行コードを空白に置換して空白を潰す
func convNewline(str, nlcode string) string {
	res := strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)

	return strings.ReplaceAll(res, " ","")
}