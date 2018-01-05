package main

// AppOpts represents required CLI options.
type AppOpts struct {
	Categories   string
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
