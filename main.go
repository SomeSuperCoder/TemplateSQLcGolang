package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/SomeSuperCoder/sqlclearning/internal/repository"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://admin:password@localhost:5432/mydatabase")
	if err != nil {
		os.Exit(0)
	}
	defer conn.Close(ctx)

	repo := repository.New(conn)
	players, _ := repo.FindAllPlayers(ctx)

	log.Println(players)

	item, err := repo.InsertItem(ctx, repository.InsertItemParams{
		Name:  "Diamond",
		Value: 1000,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(item)
}
