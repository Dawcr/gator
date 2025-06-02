package main

import (
	"context"
	"fmt"
	"html"
)

const wagslaneFeedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), wagslaneFeedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}

	printRSSFeed(feed)
	return nil
}

func printRSSFeed(feed *RSSFeed) {
	escapeAndPrint(feed.Channel.Title)
	escapeAndPrint(feed.Channel.Link)
	escapeAndPrint(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		escapeAndPrint(item.Title)
		escapeAndPrint(item.Link)
		escapeAndPrint(item.Description)
		escapeAndPrint(item.PubDate)
	}
}

func escapeAndPrint(str string) {
	fmt.Println(html.UnescapeString(str))
}
