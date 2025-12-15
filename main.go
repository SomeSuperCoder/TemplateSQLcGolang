package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/SomeSuperCoder/sqlclearning/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://admin:password@localhost:5432/mydatabase")
	if err != nil {
		os.Exit(0)
	}
	defer conn.Close(ctx)

	repo := repository.New(conn)
	value, err := repo.InsertBook(ctx, repository.InsertBookParams{
		Name:   "Some book",
		Author: "Some Author",
		Price: pgtype.Numeric{
			Int:   big.NewInt(123123),
			Exp:   0,
			Valid: true,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	books, err := repo.FindAllBooks(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(books)
}
