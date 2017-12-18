package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/html"
)

const LedFlashlightsList = "https://couponsfromchina.com/2017/06/19/ultimate-flashlight-coupons-deals-list-gearbest/"

type item struct {
	No      int
	Name    string
	Type    string
	Link    string
	Price   float64
	Sale    float64
	Percent float64
	Lowest  float64
	Coupon  string
}

// type items []*item

var items []*item

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

	sort.Slice(items, func(i, j int) bool { return items[i].Sale < items[j].Sale })

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No", "Name", "Sale, $", "Type"})
	table.SetAutoWrapText(false)
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_LEFT,
	})

	for _, v := range items {
		if v.Type != "led-flashlights" {
			continue
		}
		table.Append([]string{
			fmt.Sprintf("%d", v.No),
			v.Name,
			fmt.Sprintf("%.1f", v.Sale),
			v.Type,
		})
	}
	table.Render()
}

func forEachTR(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "tr" { // row
		cells := make([]string, 0, 9)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data != "td" {
				continue
			}
			cells = forEachTD(c, cells)
		}

		if len(cells) == 0 { // skip table header, need improve
			return
		}

		if len(cells) < 6 {
			log.Fatalf("too few cells for %v", cells[0])
		}
		// fmt.Println(strings.Join(cells, ", "))
		items = append(items, makeItem(cells))
	}
}

func makeItem(c []string) *item {
	var err error
	itm := new(item)
	itm.No, err = strconv.Atoi(c[0])
	checkError(err)
	itm.Link = c[1]
	u, err := url.Parse(itm.Link)
	checkError(err)
	itm.Type = strings.Split(u.Path, "/")[1]

	itm.Name = c[2]
	itm.Price, err = strconv.ParseFloat(strings.Trim(c[3], "$"), 64)
	checkError(err)
	itm.Sale, err = strconv.ParseFloat(strings.Trim(c[4], "$"), 64)
	checkError(err)
	if strings.HasSuffix(c[5], "%") {
		dotted := strings.Replace(c[5], ",", ".", -1)
		itm.Percent, err = strconv.ParseFloat(strings.TrimRight(dotted, "%"), 64)
		checkError(err)
		itm.Lowest, _ = strconv.ParseFloat(strings.TrimLeft(c[6], "$"), 64)
	} else {
		itm.Lowest, _ = strconv.ParseFloat(strings.TrimLeft(c[5], "$"), 64)
	}
	return itm
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
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
