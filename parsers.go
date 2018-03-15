package main

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/krasoffski/gomill/node"

	"golang.org/x/net/html"
)

func parseList(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func extractNo(n *html.Node) (int, error) {
	nu, err := strconv.Atoi(n.FirstChild.Data)
	if err != nil {
		return 0, err
	}
	return nu, nil
}

func extractName(n *html.Node) (string, error) {
	textNodes := node.Traverse(n, func(td *html.Node) bool {
		return td.Type == html.TextNode && strings.TrimSpace(td.Data) != "" && td.Data != "span"
	}, nil)
	parts := make([]string, 0, len(textNodes))
	for _, node := range textNodes {
		parts = append(parts, strings.TrimSpace(node.Data))
	}
	return strings.Join(parts, " "), nil
}

func extractCategory(n *html.Node) (string, error) {
	link, _ := extractLink(n)
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	category := strings.Split(u.Path, "/")[1]

	return category, nil
}

func extractLink(n *html.Node) (string, error) {
	link, ok := node.Attr(n.FirstChild, "href")
	if !ok {
		return "", fmt.Errorf("unable to find link node")
	}
	return link, nil
}

func extractUsualPrice(td *html.Node) (float64, error) {
	priceNode := node.Find(td, func(n *html.Node) bool {
		return strings.HasPrefix(n.Data, "$") && len(n.Data) > 1 // FIXME
	})
	if priceNode == nil {
		return 0, fmt.Errorf("unable to find price node")
	}

	priceValue, err := strconv.ParseFloat(strings.Trim(priceNode.Data, "$"), 64)
	if err != nil {
		return 0, err
	}
	return priceValue, nil
}

func extractSalePrice(n *html.Node) (float64, error) {
	textNodes := node.Traverse(n, func(td *html.Node) bool {
		return td.Type == html.TextNode && strings.TrimSpace(td.Data) != "" && !strings.HasSuffix(td.Data, "%")
	}, nil)
	var priceStr string
	for _, node := range textNodes {
		priceStr += node.Data
	}
	priceStr = strings.Split(priceStr, "\n")[0]
	priceStr = strings.Trim(strings.TrimSpace(priceStr), "$")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func extractDiscountPersent(n *html.Node) (float64, error) {
	textNodes := node.Traverse(n, func(td *html.Node) bool {
		return td.Type == html.TextNode && strings.TrimSpace(td.Data) != "" && !strings.HasSuffix(td.Data, "%")
	}, nil)
	var discountStr string
	for _, node := range textNodes {
		discountStr += node.Data
	}
	items := strings.Split(discountStr, "\n")
	if len(items) == 1 {
		return 0, nil
	}
	discountStr = strings.Trim(strings.TrimSpace(items[1]), "–%$")
	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		return 0, err
	}
	return discount, nil
}

func extractLowestPrice(n *html.Node) (float64, error) {
	textNodes := node.Traverse(n, func(td *html.Node) bool {
		return (td.Parent.Data == "span" && td.Parent.NextSibling == nil &&
			td.Type == html.TextNode && strings.TrimSpace(td.Data) != "")
	}, nil)
	if len(textNodes) == 0 {
		return 0, nil
	}
	lowestStr := strings.TrimSpace(textNodes[0].Data)
	lowest, err := strconv.ParseFloat(lowestStr, 64)
	if err != nil {
		return 0, err
	}
	return math.Abs(lowest), nil
}
