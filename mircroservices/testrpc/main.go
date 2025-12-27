package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/SomeSuperCoder/sqlclearning/internal/repository"
	"github.com/SomeSuperCoder/sqlclearning/mircroservices/rpcservices"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/jackc/pgx/v5"
)

const port = 8099

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://admin:password@localhost:5432/mydatabase")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	repo := repository.New(conn)

	mathService := new(rpcservices.MathService)
	s.RegisterService(mathService, "Math")

	bookService := &rpcservices.BookService{Repo: repo}
	s.RegisterService(bookService, "Book")

	http.Handle("/rpc", s)

	log.Printf("RPC started and is listening on :%v", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}
