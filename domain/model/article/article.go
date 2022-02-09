package domain

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ResArticle return front data
// type ArticleInfo struct {
// 	ID          uint   `gorm:"primarykey" gorm:"AUTO_INCREMENT"`
// 	Userid      string `json:"user_id"`
// 	Title       string `json:"title"`
// 	URL         string `json:"url"`
// 	Description string `json:"description"`
// 	Header      string `json:"header"`
// 	Body        string `json:"body"`
// }

type ArticleModel struct{}

type ArticleInfo struct {
	UserId       string
	Title        string
	Url          string
	Description  string
	Header       string
	ThumbnailURL string
	Body         string
}

type ArticleClientRequest struct {
	Userid    string `json:"user_id"`
	AuthToken string `json:"authtoken"`
	URL       string `json:"url"`
}

type articleOGPScanResult struct {
	URL          string
	Title        string
	Description  string
	ThumbnailURL string
	Header       string
	Body         string
}

// func UserArticleLists(userID string) ([]ArticleInfo, error) {
// 	articles, err := ArticleRepository.GetByID(userID)
// 	return articles, err
// }

func CreateArticleInfo(req ArticleClientRequest) ArticleInfo {
	article := scanArticle(req.URL)
	fmt.Println(article)

	rtnScanResult := &ArticleInfo{UserId: req.Userid, Title: article.Title, Url: article.URL, Description: article.Description, Header: article.Header, ThumbnailURL: article.ThumbnailURL, Body: article.Body}

	fmt.Println(rtnScanResult)

	return *rtnScanResult
}

// ScanArticle is scan article
func scanArticle(url string) *articleOGPScanResult {
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

	// 実装参考 https://qiita.com/Yaruki00/items/b50e346551690b158a79
	doc.Find("meta").Each(func(index int, item *goquery.Selection) {
		// Attr("rtn true", "rtn false")
		if item.AttrOr("property", "") == "og:description" {
			metaDescription = item.AttrOr("content", "")
		}

		if item.AttrOr("property", "") == "og:image" {
			metaImageURL = item.AttrOr("content", "")
		}

		if item.AttrOr("property", "") == "og:url" {
			metaURL = item.AttrOr("content", "")

			if metaURL == "" {
				metaURL = url
			}

		}
	})

	fmt.Printf("Page Title: '%s'\n", pageTitle)
	fmt.Printf("Meta Description: '%s'\n", metaDescription)
	fmt.Printf("Meta Image: '%s'\n", metaImageURL)
	fmt.Printf("Meta URL: '%s'\n", metaURL)

	pageTitle = doc.Find("title").Contents().Text()

	// 以下の指定で記事本文の取得が可能
	pageHeader := doc.Find("header").Contents().Text()
	cnvPageHeader := convNewline(pageHeader, " ")

	pageBody := doc.Find("body").Contents().Text()
	cnvPageBody := convNewline(pageBody, " ")

	pageMetaDescription := convNewline(metaDescription, " ")

	// OGP解析で取得したメタ情報
	rtnVal := articleOGPScanResult{
		URL:          metaURL,
		Title:        pageTitle,
		Description:  pageMetaDescription,
		ThumbnailURL: metaImageURL,
		Header:       cnvPageHeader,
		Body:         cnvPageBody,
	}

	fmt.Println("構造解析処理")
	fmt.Println(rtnVal.URL)

	return &rtnVal
}

// convNewline 全ての改行コードを空白に置換して空白を潰す
func convNewline(str, nlcode string) string {
	res := strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)

	return strings.ReplaceAll(res, " ", "")
}
