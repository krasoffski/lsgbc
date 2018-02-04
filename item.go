package main

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
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
	Coupon   string
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

func skip() func(n *html.Node) bool {
	return node.AnyFnN(
		func(n *html.Node) bool { return strings.TrimSpace(n.Data) == "" },
		func(n *html.Node) bool { return n.Data == "br" },
		func(n *html.Node) bool { return n.Data == "td" },
		func(n *html.Node) bool { return n.Data == "img" },
		func(n *html.Node) bool { return n.Data == "span" },
		func(n *html.Node) bool { return n.Data == "strong" },
		func(n *html.Node) bool { return n.FirstChild != nil && n.FirstChild.Data == "img" },
	)
}

func makeItemsFromURL(url string) ([]*item, error) {
	doc, err := parseList(url)
	if err != nil {
		return []*item{}, err
	}

	tableNode := node.Find(doc, func(n *html.Node) bool {
		return checkNodeID(n, "table", "alphabetically")
	})

	tbodyNode := node.Find(tableNode, func(n *html.Node) bool {
		return checkNodeID(n, "tbody", "")
	})

	trNodes := getChildren(tbodyNode, "tr")
	items := make([]*item, 0, len(trNodes))

	for _, tr := range trNodes[1:] {
		cells := node.Traverse(tr, skip(), nil)

		values, err := cellsToStrings(cells)
		if err != nil {
			return nil, err
		}
		newItem, err := makeItem(values)
		if err != nil {
			return nil, err
		}

		items = append(items, newItem)
	}
	return items, nil
}

func makeItem(c []string) (*item, error) {
	var err error
	itm := new(item)

	itm.No, err = strconv.Atoi(c[0])
	if err != nil {
		return nil, err
	}

	itm.Link = c[1]

	u, err := url.Parse(itm.Link)
	if err != nil {
		return nil, err
	}
	itm.Category = strings.Split(u.Path, "/")[1]

	itm.Name = c[2]

	itm.Usual, err = strconv.ParseFloat(strings.Trim(c[3], "$"), 64)
	if err != nil {
		return nil, err
	}

	itm.Price, err = strconv.ParseFloat(strings.Trim(c[4], "$"), 64)
	if err != nil {
		return nil, err
	}

	if strings.HasSuffix(c[5], "%") {
		dotted := strings.Replace(c[5], ",", ".", -1)
		val, err := strconv.ParseFloat(strings.TrimRight(dotted, "%"), 64)
		if err != nil {
			return nil, err
		}

		itm.Discount = math.Abs(val)
	}
	itm.Lowest, _ = strconv.ParseFloat(strings.TrimLeft(c[6], "$"), 64)
	return itm, nil
}
