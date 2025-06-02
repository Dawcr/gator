package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dawcr/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		lim, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("error parsing int: %v", err)
		}
		limit = lim
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %v", err)
	}

	printPosts(posts)
	return nil
}

func printPosts(posts []database.Post) {
	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description.String)
		fmt.Println()
	}
}
