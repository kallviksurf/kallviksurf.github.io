package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// load the html file
	var file = "test.html"
	r, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}

	var title = ""
	var header = ""
	var content = ""
	var signature = ""
	var timestamp time.Time

	// Find the review items
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title = s.Text()
		fmt.Printf("title %d: %s\n", i, title)
	})

	// header
	doc.Find("#sites-page-title-header").Each(func(i int, h *goquery.Selection) {
		header = h.Text()
	})

	// timestamp
	doc.Find("span.announcementsPostTimestamp span").Each(func(i int, s *goquery.Selection) {

		// For each item found, get the title
		txt := s.Text()
		if txt != "" {
			parts := strings.Split(txt, " ")
			// for ii := 0; ii < len(parts); ii++ {
			// 	fmt.Printf("timestamp part %d: %s\n", ii, parts[ii])
			// }

			txt = fmt.Sprintf("%s %s", parts[0], parts[2])
			timestamp, err = time.Parse("2.1.2006 15.04", txt)
			if err != nil {
				fmt.Printf("could not parse time from %s, %s\n", err, txt)
			}

		}
		fmt.Printf("timestamp %d: %s\n", i, timestamp)
	})
	doc.Find("#sites-canvas-main-content td.sites-layout-tile > div > div").Contents().Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "#text" {
			trimmed := strings.Trim(s.Text(), " \n")
			if trimmed != "" {
				fmt.Printf(">>> (%d) >>> '%s'\n", i, trimmed)
				content = content + trimmed
			}
		}
	})

	doc.Find("#sites-canvas-main-content td.sites-layout-tile > div").Contents().Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "#text" {
			trimmed := strings.Trim(s.Text(), " \n")
			if trimmed != "" {
				signature = trimmed
			}
		}
	})

	fmt.Printf("header: %s\n", header)
	fmt.Printf("timestamp: %s\n", timestamp)
	fmt.Printf("title: %s\n", title)
	fmt.Printf("content: %s\n", content)
	fmt.Printf("signature: %s\n", signature)

	var md = fmt.Sprintf(`+++
title="%s"
date="%s"
+++

# %s

%s

%s
`, header, timestamp.Format(time.RFC3339), header, content, signature)

	fmt.Println(md)

}
