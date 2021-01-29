package gomob

import (
	"context"
	"log"
	"net/url"
	"path/filepath"
	"regexp"

	firebase "firebase.google.com/go"
	"github.com/PuerkitoBio/goquery"
	"google.golang.org/api/option"
)

// ProductField is Product document field structure.
type ProductField struct {
	Url        string
	Id         string
	CircleName string
	CV         []string
	Genre      []string
	Title      string
}

type ProductInfo struct {
	Field ProductField
}

var AllowDomain = [...]string{"dlsite.com"}

func SetProductInfo(product ProductInfo) bool {
	result := true

	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase-setting.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Collection("products").Doc(product.Field.Id).Set(ctx, product.Field)
	if err != nil {
		log.Fatal(err)
		result = false
	}
	return result
}

func GetProductCode(urlstr string) string {
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}
	rep := regexp.MustCompile(`.html.*$`)
	code := filepath.Base(rep.ReplaceAllString(u.Path, ""))
	//TODO: プロダクトコードのチェック
	return code
}

func GetProductInfo(url string) ProductInfo {
	var result ProductInfo
	// URLをアフィリエイトリンクに書き換える
	result.Field.Url = url
	result.Field.Id = GetProductCode(url)
	//Add fields
	//TODO: 意図しないサイトへのアクセスをさせないようにする
	var genres []string
	var actors []string
	doc, _ := goquery.NewDocument(url)
	result.Field.CircleName = doc.Find("#work_maker > tbody > tr > td > span > a").Text()
	result.Field.Title = doc.Find("#work_name > a").Text()
	doc.Find("div.main_genre > a").Each(func(_ int, s *goquery.Selection) {
		genres = append(genres, s.Text())
	})
	result.Field.Genre = genres
	doc.Find("#work_outline > tbody > tr > th:contains('声優')").Parent().Find("td > a").Each(func(_ int, s *goquery.Selection) {
		actors = append(actors, s.Text())
	})
	result.Field.CV = actors
	/*
		doc.Find("#work_outline > tbody > tr > th:contains('声優')").Parent().Find("td > a").Each(func(_ int, s *goquery.Selection) {
			fmt.Println(s.Text())
		})
	*/

	return result
}
