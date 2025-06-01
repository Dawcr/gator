package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/dawcr/gator/internal/config"
	"github.com/dawcr/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState := &state{
		cfg: &cfg,
	}

	db, err := sql.Open("postgres", appState.cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	appState.db = database.New(db)

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		log.Fatal("usage: cli <command> [args...]")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.run(appState, cmd); err != nil {
		log.Fatalf("error running command %v: %v", cmd.Name, err)
	}
}
