package pager

import "math"

// Page ...
type Page struct {
	Index int
}

// Pages ...
type Pages struct {
	CurrentPage        int
	Pages              []Page
	PageCount          int
	ItemsPerPage       int
	ItemsLength        int
	CurrentOffset      int
	HasNext, HasPrev   bool
	PrevPage, NextPage Page
}

func calculateOffset(p *Pages) int {
	switch {
	case p.CurrentPage > p.PageCount:
		p.CurrentPage = p.PageCount
	case p.CurrentPage < 1:
		p.CurrentPage = 1
	}
	o := (p.CurrentPage - 1) * p.ItemsPerPage
	switch {
	case o > p.ItemsLength:
		p.CurrentPage = p.PageCount
		o = calculateOffset(p)
	case o < 0:
		o = 0
		p.CurrentPage = 1
	}
	return o
}

// New ...
func New(current, itemsPerPage, itemsLength int) Pages {
	p := Pages{}

	p.CurrentPage = current
	p.ItemsPerPage = itemsPerPage
	p.ItemsLength = itemsLength

	d := float64(p.ItemsLength) / float64(p.ItemsPerPage)
	p.PageCount = int(math.Ceil(d))

	p.CurrentOffset = calculateOffset(&p)
	p.HasPrev = p.CurrentPage > 1
	p.HasNext = p.CurrentPage < p.PageCount

	switch {
	case p.HasNext:
		p.NextPage = Page{p.CurrentPage + 1}
	case p.HasPrev:
		p.PrevPage = Page{p.CurrentPage - 1}
	}

	for i := 1; i <= p.PageCount; i++ {
		p.Pages = append(p.Pages, Page{i})
	}

	return p
}

// Range return []Page within the given range
func (p Pages) Range(min, max int) []Page {
	pages := []Page{}

	for i := min; i <= max; i++ {
		for _, page := range p.Pages {
			if page.Index == i {
				pages = append(pages, page)
			}
		}
	}

	return pages
}

// Margin return []Page within a given margin from the CurrentPage
func (p Pages) Margin(m int) []Page {
	return p.Range(p.CurrentPage-m, p.CurrentPage+m)
}
