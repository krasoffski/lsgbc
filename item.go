package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const LedFlashlightsList = "https://couponsfromchina.com/2017/06/19/ultimate-flashlight-coupons-deals-list-gearbest/"

type item struct {
	No       int
	Name     string
	Link     string
	Price    int
	Sale     int
	Discount int
	Lowest   int
	Coupon   string
}

type items []item

func main() {
	resp, err := http.Get(LedFlashlightsList)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	forEachNode(doc, forEachTR, nil)
}

func forEachTR(n *html.Node) {
	items := make([]string, 0)
	if n.Type == html.ElementNode && n.Data == "tr" { // row
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			item := forEachTD(c)
			if len(item) < 1 {
				continue
			}
			items = append(items, item...)
		}
		fmt.Println(strings.Join(items, ", "))
	}
}

func forEachTD(n *html.Node) []string {
	// if n.Type == html.ElementNode && n.Data != "td" {
	// 	return nil
	// }
	items := make([]string, 0)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.TextNode:
			items = append(items, strings.TrimSpace(c.Data))
		case html.ElementNode:
			if c.Data != "a" {
				continue
			}
			for _, a := range c.Attr {
				if a.Key != "href" {
					continue
				}
				// items = append(items, a.Val)
			}
			qq := c.FirstChild
			items = append(items, qq.Data)
		}
	}
	return items
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
