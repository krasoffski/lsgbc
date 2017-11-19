package lsgbc

type item struct {
	No     int
	Name   string
	Link   string
	Price  int
	Sale   int
	Lowest int
	Coupon string
}

type items []item
