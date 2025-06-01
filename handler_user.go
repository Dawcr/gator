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
		return fmt.Errorf("usage: cli %s <name>", cmd.Name)
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
		return fmt.Errorf("usage: cli %s <name>", cmd.Name)
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

	fmt.Printf("user %s has been created, %v\n", usr.Name, usr)
	return nil
}
