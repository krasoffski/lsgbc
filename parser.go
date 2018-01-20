package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

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

func getChildren(n *html.Node, data string) []*html.Node {
	nodes := make([]*html.Node, 0)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == data {
			nodes = append(nodes, c)
		}
	}
	return nodes
}

func getAttrByName(n *html.Node, attr string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val, true
		}
	}
	return "", false
}

func cellsToStrings(cells []*html.Node) ([]string, error) {
	nLen := len(cells)
	if nLen < 6 {
		return nil, fmt.Errorf("invalid number of nodes %d", nLen)
	}
	arr := make([]string, 7)
	arr[0] = cells[0].Data

	var ok bool
	if arr[1], ok = getAttrByName(cells[1], "href"); !ok {
		return nil, errors.New("unable to found item link")
	}

	arr[2] = cells[2].Data
	arr[3] = cells[3].Data
	arr[4] = cells[3].Data

	if len(cells) == 6 {
		arr[6] = cells[5].Data
	} else {
		arr[5] = strings.TrimSpace(cells[5].Data)
		arr[6] = cells[6].Data
	}
	return arr, nil
}

func getTrCells(n *html.Node) []string {
	cells := make([]string, 0)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.TextNode:
			cells = append(cells, strings.TrimSpace(c.Data))
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
				cells = append(cells, url, childNode.Data)
			case "span":
				cells = append(cells, c.FirstChild.FirstChild.Data) // NOTE: OMG
			}
		}
	}
	fmt.Println(cells)
	return cells
}

func findNode(n *html.Node, check func(*html.Node) bool) *html.Node {
	if check(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := findNode(c, check)
		if node != nil {
			return node
		}
	}

	return nil
}

func traverseNode(n *html.Node, skip func(*html.Node) bool, nodes []*html.Node) []*html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !skip(c) {
			nodes = append(nodes, c)
		}
		nodes = traverseNode(c, skip, nodes)
	}
	return nodes
}
