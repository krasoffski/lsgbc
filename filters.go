package main

func sortOut(items []*item, opts *AppOpts) []*item {

	filtered := make([]*item, 0)

	for _, v := range items {
		if !globWords(v.Category, uniqOpts(opts.Categories)) {
			continue
		}

		if !globWords(v.Name, uniqOpts(opts.Names)) {
			continue
		}

		if v.Price < opts.MinPrice || v.Price > opts.MaxPrice {
			continue
		}

		if opts.ShowBest {
			if v.Price > v.Lowest*1.1 {
				continue
			}
		}
		filtered = append(filtered, v)
	}

	return filtered
}
