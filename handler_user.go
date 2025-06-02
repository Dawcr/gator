package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dawcr/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	usr, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user %s does not exist in database", cmd.Args[0])
	}

	if err := s.cfg.SetUser(usr.Name); err != nil {
		return fmt.Errorf("couldn't switch current user to %v", err)
	}

	fmt.Printf("username has been set to %v", usr.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	usr, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	})

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	if err := s.cfg.SetUser(usr.Name); err != nil {
		return err
	}

	fmt.Printf("user %s has been created\n", usr.Name)
	printUser(usr)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving users list from db: %v", err)
	}
	for _, user := range users {
		fmt.Printf("* %v", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Printf("\n")
	}
	return nil
}
