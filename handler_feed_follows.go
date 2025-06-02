package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dawcr/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedFromURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't retrieve info for url: %v", err)
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error while creating feed follow: %v", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feed_follow.UserName, feed_follow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, usr database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), usr.ID)
	if err != nil {
		return fmt.Errorf("error retrieving feed follows for current user: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Println("Current user does not follow any feed.")
		return nil
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

func followCreated(s *state, url string, user database.User) error {
	return handlerFollow(s, command{
		Name: "follow",
		Args: []string{url},
	},
		user)
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedFromURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error retrieving feed info: %v", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing: %v", err)
	}

	fmt.Printf("%v unfollowed\n", feed.Name)
	return nil
}
