package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.ResetDB(context.Background()); err != nil {
		return fmt.Errorf("error resetting db: %v", err)
	}
	fmt.Printf("Db reset\n")
	return nil
}
