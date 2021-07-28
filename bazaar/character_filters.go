package bazaar

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func HasAlumni(e *colly.HTMLElement) bool {
	var hasAlumni = false
	e.DOM.Find("div.AuctionBodyBlock.SpecialCharacterFeatures div.Entry").Each(func(i int, selection *goquery.Selection) {
		selection.Siblings().Each(func(i int, selection *goquery.Selection) {
			if hasAlumni == true {
				return
			}
			hasAlumni = strings.Contains(selection.Text(), "Alumni")
		})
	})

	return hasAlumni
}
