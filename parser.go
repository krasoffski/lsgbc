package main

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type collectorFunc func([]*html.Node, *html.Node) []*html.Node
type checkFunc func(*html.Node) bool

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

func checkNodeID(n *html.Node, data, id string) bool {
	if n.Type == html.ElementNode && n.Data == data {
		if id == "" {
			return true
		}
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
	}
	return false
}

func getTableTrNodes(n *html.Node) []*html.Node {
	nodes := make([]*html.Node, 0)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "tr" { // row
			nodes = append(nodes, c)
		}
	}
	return nodes
}

// NOTE: think about passing pointer to slice instead of return.
func forEachTD(n *html.Node, items []string) []string {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.TextNode:
			items = append(items, strings.TrimSpace(c.Data))
		case html.ElementNode:
			switch c.Data {
			case "a":
				var url string
				for _, a := range c.Attr {
					if a.Key != "href" {
						continue
					}
					url = a.Val
					break // ignoring all other attrs
				}
				childNode := c.FirstChild
				if childNode.Data == "img" { // ugly hack for skipping img cell
					continue
				}
				items = append(items, url, childNode.Data)
			case "span":
				items = append(items, c.FirstChild.FirstChild.Data) // NOTE: OMG
			}
		}
	}
	return items
}

func findNode(n *html.Node, cf checkFunc) *html.Node {
	if cf(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := findNode(c, cf)
		if node != nil {
			return node
		}
	}

	return nil
}
