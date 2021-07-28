package main

import (
	"io/ioutil"
	"os"
	"tibiaScrapper/bazaar"
	"tibiaScrapper/utils"
	"time"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type character struct {
	Url       string `json:"url"`
	Name      string `json:"char_name"`
	CrawledAt time.Time `json:"crawled_at"`
}

type page struct {
	Link string
	Index string
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	const baseUrl = "https://www.tibia.com/charactertrade/?subtopic=currentcharactertrades"
	var c = colly.NewCollector()

	var (
		pages             []page
		characters        []character
		visitedPagesIndex = []string{"1"}
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", utils.RandomString())
		log.Infoln("Accessing: ", r.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		log.Infoln("Visited: ", response.Request.URL)
		pages = append(
			pages,
			page{ Link:  baseUrl, Index: "1"},
		)
	})

	// Find Character Details
	c.OnHTML(".Auction", func(e *colly.HTMLElement) {
		charDetailsDOM := e.DOM.Find("div.AuctionCharacterName > a")

		if bazaar.HasAlumni(e) {
			charUrl, _ := charDetailsDOM.Attr("href")
			charName := charDetailsDOM.Text()

			log.WithFields(log.Fields{
				"url": charUrl,
				"name": charName,
			}).Warnln("Character with Alumni achievement found!!")

			characters = append(characters, character{
				charUrl,
				charName,
				time.Now(),
			})
		}
	})

	c.OnHTML(".PageNavigation span.PageLink a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		pageIndex := e.Text

		//Avoid shortcuts to Last Page and First Page to follow natural path
		if pageIndex == "Last Page" || pageIndex == "First Page" {
			return
		}

		// If we already visited the page we can skip
		if utils.InArray(visitedPagesIndex, pageIndex) {
			return
		}

		visitedPagesIndex = append(visitedPagesIndex, pageIndex)

		// Wait to don't get banned
		utils.SleepRandom(3, 1)

		//Access the next page
		visitUrl(c, e.Request.AbsoluteURL(link))
	})

	visitUrl(c, baseUrl)

	results, err := utils.ToJson(characters)

	utils.HandleErrorDefault(err)

	err = ioutil.WriteFile("outputs/"+ utils.GetFileNameWithTime("bazaar", ".json"), results, 0644)

	utils.HandleErrorDefault(err)
}

func visitUrl(c *colly.Collector, url string)  {
	err := c.Visit(url)

	utils.HandleErrorDefault(err)
}


