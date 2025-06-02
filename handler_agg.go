package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dawcr/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing time_between_reqs: %v", err)
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feedInfo, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("error getting next feed to fetch: %v", err)
		return
	}
	log.Println("found a feed to fetch")

	scrapeFeed(s.db, feedInfo)
}

func scrapeFeed(db *database.Queries, feedInfo database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feedInfo.ID)
	if err != nil {
		log.Printf("error marking feed as fetched: %v", err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feedInfo.Url)
	if err != nil {
		log.Printf("error fetching feed: %v", err)
		return
	}

	for _, feedItem := range feedData.Channel.Item {
		fmt.Printf("Found post: %v\n", feedItem.Title)
		parsedTime, err := time.Parse(time.RFC1123Z, feedItem.PubDate)
		if err != nil {
			log.Printf("error parsing time: %v", err)
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     feedItem.Title,
			Url:       feedItem.Link,
			Description: sql.NullString{
				String: feedItem.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  parsedTime,
				Valid: true,
			},
			FeedID: feedInfo.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feedInfo.Name, len(feedData.Channel.Item))
}
