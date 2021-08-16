package otherpackage

import "fmt"

import (
	"github.com/gocolly/colly"
)

func BaiduSpider() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.baidu.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".result.c-container", func(e *colly.HTMLElement) {
		//link := e.Attr("href")
		//// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
		fmt.Println(e.ChildText("h3.t > a"))
		fmt.Println(e.ChildAttr("h3.t > a", "href"))
	})

	c.OnHTML("a.n", func(e *colly.HTMLElement) {
		_ = c.Visit("https://www.baidu.com" + e.Attr("href"))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&rsv_idx=1&tn=baidu&wd=%E6%97%AD%E9%98%B3&fenlei=256&oq=%25E6%2597%25AD%25E9%2598%25B3&rsv_pq=db3ec12c0007f413&rsv_t=3959GBJI3j%2B5OlWk5ojMXhvTUWwaxsIpOjaqSgoa2T%2FTHUhtv1vhCCV1jIc&rqlang=cn&rsv_enter=0&rsv_dl=tb&rsv_btype=t")
}
	