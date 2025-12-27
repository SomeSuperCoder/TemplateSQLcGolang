package rpcservices

import (
	"net/http"

	"github.com/SomeSuperCoder/sqlclearning/internal/repository"
)

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

func (b *BookService) Update(r *http.Request, args *repository.UpdateBookParams, reply *repository.Book) error {
	book, err := b.Repo.UpdateBook(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = book

	return nil
}
