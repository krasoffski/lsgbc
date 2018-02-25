package main

import (
	"fmt"
	"strings"

	"github.com/krasoffski/gomill/node"
	glob "github.com/ryanuber/go-glob"
	"golang.org/x/net/html"
)

var skipNodeData = map[string]bool{
	"":       true,
	"br":     true,
	"td":     true,
	"img":    true,
	"strong": true,
	"span":   true,
}

type item struct {
	No       int
	Name     string
	Category string
	Link     string
	Usual    float64
	Price    float64
	Discount float64
	Lowest   float64
}

func globWords(subj string, patterns map[string]struct{}) bool {
	subjLower := strings.ToLower(subj)
	for p := range patterns {
		p = strings.ToLower(strings.Trim(p, "*"))
		if glob.Glob(p+"*", subjLower) {
			return true
		}
	}
	return false
}

func uniqOpts(s string) map[string]struct{} {
	unique := make(map[string]struct{})
	for _, w := range strings.Split(s, ",") {
		unique[w] = struct{}{}
	}
	return unique
}

func nonZero(val float64) string {
	var printable string
	if val > 0.0 {
		printable = fmt.Sprintf("%.1f", val)
	} else {
		printable = "-"
	}
	return printable
}

func makeItemsFromURL(url string) ([]*item, error) {
	doc, err := parseList(url)
	if err != nil {
		return []*item{}, err
	}

	tableNode := node.Find(doc, func(n *html.Node) bool {
		id, ok := node.Attr(n, "id")
		return n.Type == html.ElementNode && n.Data == "table" && ok && id == "alphabetically"
	})

	tbodyNode := node.Find(tableNode, func(n *html.Node) bool {
		_, ok := node.Attr(n, "id")
		return n.Type == html.ElementNode && n.Data == "tbody" && !ok
	})

	trNodes := node.Children(tbodyNode, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "tr"
	})

	items := make([]*item, 0, len(trNodes))

	for _, tr := range trNodes[1:] {
		itm, err := makeItem(tr)
		if err != nil {
			return nil, err
		}
		items = append(items, itm)
	}
	return items, nil
}

func makeItem(n *html.Node) (*item, error) {
	tdNodes := node.Children(n, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "td"
	})
	var err error
	itm := new(item)
	if itm.No, err = extractNo(tdNodes[0]); err != nil {
		return nil, fmt.Errorf("unable to parse No %v", err)
	}

	if itm.Name, err = extractName(tdNodes[2]); err != nil {
		return nil, fmt.Errorf("unable to parse Name %v", err)
	}

	if itm.Category, err = extractCategory(tdNodes[2]); err != nil {
		return nil, fmt.Errorf("unable to parse Category %v", err)
	}

	if itm.Link, err = extractLink(tdNodes[2]); err != nil {
		return nil, fmt.Errorf("unable to parse Link %v", err)
	}

	if itm.Usual, err = extractUsualPrice(tdNodes[3]); err != nil {
		return nil, fmt.Errorf("unable to parse Usual price %v", err)
	}

	if itm.Price, err = extractSalePrice(tdNodes[4]); err != nil {
		return nil, fmt.Errorf("unable to parse Sale price %v", err)
	}

	if itm.Discount, err = extractDiscountPersent(tdNodes[4]); err != nil {
		return nil, fmt.Errorf("unable to parse Discount persent %v", err)
	}

	if itm.Lowest, err = extractLowestPrice(tdNodes[5]); err != nil {
		return nil, fmt.Errorf("unable to parse Lowest price %v", err)
	}

	return itm, nil
}
