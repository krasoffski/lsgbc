package main

// AppOpts represents required CLI options.
type AppOpts struct {
	CompactTable bool
	FlashSale    bool
	List         string
	MaxPrice     float64
	MinPrice     float64
	Names        string
	ShowBest     bool
	SortBy       string
	Version      bool
}
