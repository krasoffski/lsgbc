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
	return n.FirstChild.FirstChild.Data, nil
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
		return "", fmt.Errorf("unable to find link")
	}
	return link, nil
}

func extractUsualPrice(td *html.Node) (float64, error) {
	priceNode := node.Find(td, func(n *html.Node) bool {
		return n.Parent == td && strings.HasPrefix(n.Data, "$")
	})
	if priceNode == nil {
		return 0, fmt.Errorf("Unable to find price")
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
	for _, node := range textNodes[5:] {
		priceStr += node.Data
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func extractDiscountPersent(n *html.Node) (float64, error) {
	textNodes := node.Traverse(n, func(td *html.Node) bool {
		return td.Type == html.TextNode && strings.TrimSpace(td.Data) != "" && strings.HasSuffix(td.Data, "%")
	}, nil)
	if len(textNodes) == 0 {
		return 0, nil
	}
	discountStr := strings.Trim(strings.TrimSpace(textNodes[0].Data), "%")
	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		return 0, err
	}
	return math.Abs(discount), nil
}

func extractLowestPrice(n *html.Node) (float64, error) {
	textNodes := node.Traverse(n, func(td *html.Node) bool {
		return td.Parent == n && td.Type == html.TextNode && strings.TrimSpace(td.Data) != "" && strings.HasPrefix(td.Data, "$")
	}, nil)
	if len(textNodes) == 0 {
		return 0, nil
	}
	lowestStr := strings.Trim(strings.TrimSpace(textNodes[0].Data), "$")
	lowest, _ := strconv.ParseFloat(lowestStr, 64)
	return math.Abs(lowest), nil
}
