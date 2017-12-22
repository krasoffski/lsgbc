package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ogier/pflag"
	"github.com/olekukonko/tablewriter"
	glob "github.com/ryanuber/go-glob"
	"golang.org/x/net/html"
)

// Version is version of package.
var Version = "0.0.0"

const ledFlashlightsList = "https://couponsfromchina.com/2017/06/19/ultimate-flashlight-coupons-deals-list-gearbest/"

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

// type items []*item

var items []*item

func main() {
	minPrice := pflag.Float64P("min-price", "m", 0.0, "minimal discount price")
	maxPrice := pflag.Float64P("max-price", "M", 1000.0, "maximum discount price")
	categoriesGlob := pflag.StringP("categories", "c", "*",
		"comma separated list of categories matching with glob pattern (case sensitive)")
	namesGlob := pflag.StringP("names", "n", "*",
		"comma separated list of names matching with glob pattern (case sensitive)")
	compactTable := pflag.BoolP("compact", "C", false, "use compact table representation")
	version := pflag.BoolP("version", "V", false, "show version and exit")
	onlyBest := pflag.Bool("best", false, "show only best deals")
	_ = pflag.StringP("sort-by", "S", "price", "not yet implemented")
	_ = pflag.StringP("list", "l", "flashlight", "not yet implemented")
	_ = pflag.BoolP("deskending", "d", false, "not yet implemented")
	pflag.Parse()

	if *version {
		fmt.Printf("Version is: %s\n", Version)
		os.Exit(0)
	}

	categories := uniqOpts(*categoriesGlob)
	names := uniqOpts(*namesGlob)

	doc, err := parseList(ledFlashlightsList)
	if err != nil {
		log.Fatalln(err)
	}

	forEachNode(doc, forEachTR, nil)

	sort.Slice(items, func(i, j int) bool { return items[i].Price < items[j].Price })

	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetBorder(false)

	if *compactTable {
		table.SetHeader([]string{"#", "N", "P, $", "D, %", "L, $"})
	} else {
		table.SetHeader([]string{"Nu", "Name", "Price, $", "Discount, %", "Lowest, $", "Category"})
	}

	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		// tablewriter.ALIGN_LEFT,
	})

	var count int

	for _, v := range items {
		// NOTE: add filters here.
		if !globWords(v.Category, categories) {
			continue
		}

		if !globWords(v.Name, names) {
			continue
		}

		if v.Price < *minPrice || v.Price > *maxPrice {
			continue
		}

		if *onlyBest {
			if v.Price > v.Lowest*1.1 {
				continue
			}
		}

		row := make([]string, 0, 6) // Do not allocate new memory during append.

		row = append(row,
			fmt.Sprintf("%d", v.No),
			v.Name,
			fmt.Sprintf("%.1f", v.Price),
			nonZero(v.Discount),
			nonZero(v.Lowest),
		)
		if !*compactTable {
			row = append(row, v.Category)
		}
		table.Append(row)
		count++
	}

	if *compactTable {
		table.SetFooter([]string{"", "", "", "", fmt.Sprintf("%d", count)})
	} else {
		table.SetFooter([]string{"", "", "", "", "Items", fmt.Sprintf("%d", count)})
	}

	table.Render()
}

func parseList(url string) (*html.Node, error) {
	resp, err := http.Get(ledFlashlightsList)
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

func globWords(s string, words map[string]struct{}) bool {
	for c := range words {
		if glob.Glob(c, s) {
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
	itm.Category = strings.Split(u.Path, "/")[1]

	itm.Name = c[2]
	itm.Usual, err = strconv.ParseFloat(strings.Trim(c[3], "$"), 64)
	checkError(err)
	itm.Price, err = strconv.ParseFloat(strings.Trim(c[4], "$"), 64)
	checkError(err)
	if strings.HasSuffix(c[5], "%") {
		dotted := strings.Replace(c[5], ",", ".", -1)
		val, err := strconv.ParseFloat(strings.TrimRight(dotted, "%"), 64)
		checkError(err)
		itm.Discount = math.Abs(val)
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
