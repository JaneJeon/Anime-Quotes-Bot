package quote

import (
	"fmt"
	
	"github.com/PuerkitoBio/goquery"
)

type Brainyquote quotes

const baseUrlForm = "https://www.brainyquote.com/topics/%s?SPvm=1&vm=l"
const urlForm = "https://www.brainyquote.com/topics/%s_%d?vm=l"

func New(topics []string) (*Brainyquote, error) {
	return nil, nil
}

func getAllQuotes(topic string) ([]Quote, error) {
	doc, err := goquery.NewDocument(fmt.Sprintf(baseUrlForm, topic))
	if err != nil {
		return nil, err
	}
	
	links := doc.Find("a")
	var result []Quote
	
	for i := range links.Nodes {
		node := links.Eq(i)
		if attr, _ := node.Attr("title"); attr == "view quote" {
			q := Quote{node.Text(), node.Next().Text()}
			result = appendSet(q, result)
		}
	}
	
	return result, nil
}

func (*Brainyquote) GetRandomQuote() (Quote, error) {
	return Quote{}, nil
}

func (*Brainyquote) refresh() error {
	return nil
}