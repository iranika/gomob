package gomob

import (
	"log"
	"net/url"
	"path/filepath"
	"regexp"

	"cloud.google.com/go/firestore"
)

//TODO: Write commnent AnswerLimit
type AnswerLimit struct {
	Circle uint
	CV     uint
}

type Answer string

// DeckField is decks document field structure.
type DeckField struct {
	Date       string
	Answers    []Answer
	AnswerType AnswerLimit
	Auther     string
	Hints      []string
	Title      string
}

//TODO: Impl function
func DeckGetProductInfo(firestore *firestore.Client, urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	rep := regexp.MustCompile(`.html$`)
	code := filepath.Base(rep.ReplaceAllString(u.Path, ""))
	//TODO: プロダクトコードのチェック
	return code
}
