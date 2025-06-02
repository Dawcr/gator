package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dawcr/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("missing follow target")
	}

	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't retrieve info for current user: %v", err)
	}

	feed, err := s.db.GetFeedFromURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't retrieve info for url: %v", err)
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    usr.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error while creating feed follow: %v", err)
	}

	fmt.Printf("%v successfully followed %v", feed_follow.UserName, feed_follow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving feed follows for current user: %v", err)
	}

	printFeedFollows(feeds)

	return nil
}

func printFeedFollows(feeds []database.GetFeedFollowsForUserRow) {
	fmt.Printf("currently following: \n")
	for _, feed := range feeds {
		fmt.Printf("- %v\n", feed.FeedName)
	}
}

func followCreated(s *state, url string) error {
	return handlerFollow(s, command{
		Name: "follow",
		Args: []string{url},
	})
}
