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
	forEachNode(doc, printTR, nil)
}

func printTR(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "td" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				fmt.Printf("%s ", strings.TrimSpace(c.Data))
			}
			fmt.Println()
		}

	}

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
