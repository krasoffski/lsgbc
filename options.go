package main

// Options represents required CLI options.
type Options struct {
	CompactTable bool
	FlashSale    bool
	List         string
	MaxPrice     float64
	MinPrice     float64
	Names        string
	ShowBest     bool
	SortBy       string
	Version      bool
	Verbose      bool
}
