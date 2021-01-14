package gomob

import (
	"log"
	"net/url"
	"path/filepath"
	"regexp"

	"github.com/PuerkitoBio/goquery"
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

func getProductCode(urlstr string) string {
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}
	rep := regexp.MustCompile(`.html$`)
	code := filepath.Base(rep.ReplaceAllString(u.Path, ""))
	//TODO: プロダクトコードのチェック
	return code
}

func getProductInfo(url string) ProductInfo {
	var result ProductInfo
	result.Field.Url = url
	result.Field.Id = getProductCode(url)
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
