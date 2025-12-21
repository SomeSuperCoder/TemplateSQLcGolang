package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/SomeSuperCoder/sqlclearning/internal/repository"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/jackc/pgx/v5"
)

const port = 8099

type Args struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Reply struct {
	Result float64 `json:"result"`
}

type MathService struct{}

func (m *MathService) Add(r *http.Request, args *Args, reply *Reply) error {
	reply.Result = args.A + args.B
	return nil
}

func (m *MathService) Divide(r *http.Request, args *Args, reply *Reply) error {
	if args.B == 0 {
		return fmt.Errorf("division by zero")
	}
	reply.Result = args.A / args.B
	return nil
}

type BookService struct {
	Repo *repository.Queries
}

func (b *BookService) FindAll(r *http.Request, args *any, reply *[]repository.Book) error {
	books, err := b.Repo.FindAllBooks(r.Context())
	if err != nil {
		return err
	}
	*reply = books
	return nil
}

func (b *BookService) Insert(r *http.Request, args *repository.InsertBookParams, reply *repository.Book) error {
	book, err := b.Repo.InsertBook(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = book

	return nil
}

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

	mathService := new(MathService)
	s.RegisterService(mathService, "Math")

	bookService := &BookService{Repo: repo}
	s.RegisterService(bookService, "Book")

	http.Handle("/rpc", s)

	log.Printf("RPC started and is listening on :%v", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}
