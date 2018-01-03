package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/ogier/pflag"
)

// Version is version of package.
var Version = "0.0.0"

var items []*item

func main() {
	var opts = AppOpts{}
	// This code smells.
	allowedList := make([]string, 0, len(Links))
	for k := range Links {
		allowedList = append(allowedList, k)
	}
	sort.Slice(allowedList, func(i, j int) bool { return allowedList[i] < allowedList[j] })
	allowed := strings.Join(allowedList, ",")

	pflag.BoolVarP(&opts.CompactTable, "compact", "C", false, "use compact table representation")
	pflag.BoolVarP(&opts.ShowBest, "best", "B", false, "show only best deals")
	pflag.BoolVarP(&opts.Version, "version", "V", false, "show version and exit")
	pflag.Float64VarP(&opts.MaxPrice, "max-price", "M", 1000.0, "maximum discount price")
	pflag.Float64VarP(&opts.MinPrice, "min-price", "m", 0.0, "minimal discount price")
	pflag.StringVarP(&opts.Categories, "categories", "c", "*", "comma separated list of categories (case sensitive), e.g. 'aa,b*,cc'")
	pflag.StringVarP(&opts.List, "list", "l", "flashlight", fmt.Sprintf("used coupons list, one from: %s", allowed))
	pflag.StringVarP(&opts.Names, "names", "n", "*", "comma separated list of names (case sensitive), e.g. 'xx,y*,zz'")
	pflag.StringVarP(&opts.SortBy, "sort-by", "S", "price", "sort table by column, 'price' or 'discount'")
	pflag.Parse()

	if opts.Version {
		fmt.Printf("Version is: %s\n", Version)
		os.Exit(0)
	}

	url, ok := Links[opts.List]
	if !ok {
		log.Fatalf("invalid choice '%s', allowed one from: %s\n", opts.List, allowed)
	}

	doc, err := parseList(url)
	if err != nil {
		log.Fatalln(err)
	}

	forEachNode(doc, forEachTR, skipByPriceTable)

	filtered := sortOut(items, &opts)

	var sortByFunc func(int, int) bool
	switch opts.SortBy {
	case "d", "discount":
		sortByFunc = func(i, j int) bool { return filtered[i].Discount > filtered[j].Discount }
	case "p", "price":
		sortByFunc = func(i, j int) bool { return filtered[i].Price < filtered[j].Price }
	}
	sort.Slice(filtered, sortByFunc)
	fmt.Println()
	printfTable(os.Stdout, filtered, &opts)
}
