package main

import (
	"github.com/mmcdole/gofeed"
)

func urlToFeed(url string) (gofeed.Feed, error) {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL(url)

	return *feed, nil
}
