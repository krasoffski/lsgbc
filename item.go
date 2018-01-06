package main

import (
	"fmt"
	"log"
	"math"
	"net/url"
	"strconv"
	"strings"

	glob "github.com/ryanuber/go-glob"
)

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

// func makeItems(rows [][]string) []*item {

// }

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
