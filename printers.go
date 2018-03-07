package main

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

func printfTable(out io.Writer, lst []*item, opts *AppOpts) {
	table := tablewriter.NewWriter(out)
	table.SetAutoWrapText(false)
	table.SetBorder(false)

	if opts.CompactTable {
		table.SetHeader([]string{"#", "N", "P, $", "D, %", "L, $"})
	} else {
		table.SetHeader([]string{"Nu", "Name", "Price, $", "Discount, %", "Lowest, $"})
	}

	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
	})

	var count int
	for _, v := range lst {
		row := make([]string, 0, 6) // Do not allocate new memory during append.

		row = append(row,
			fmt.Sprintf("%d", v.No),
			v.Name,
			fmt.Sprintf("%.1f", v.Price),
			nonZero(v.Discount),
			nonZero(v.Lowest),
		)
		table.Append(row)
		count++
	}

	if opts.CompactTable {
		table.SetFooter([]string{"", "", "", "", fmt.Sprintf("%d", count)})
	} else {
		table.SetFooter([]string{"", "", "", "Items", fmt.Sprintf("%d", count)})
	}

	table.Render()
}
