package rpcservices

import (
	"net/http"

	"github.com/SomeSuperCoder/sqlclearning/internal/repository"
)

type BookService struct {
	Repo *repository.Queries
}

func (s *BookService) FindAll(r *http.Request, args *any, reply *[]repository.Book) error {
	books, err := s.Repo.FindAllBooks(r.Context())
	if err != nil {
		return err
	}
	*reply = books

	return nil
}

func (s *BookService) Insert(r *http.Request, args *repository.InsertBookParams, reply *repository.Book) error {
	book, err := s.Repo.InsertBook(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = book

	return nil
}

func (s *BookService) Update(r *http.Request, args *repository.UpdateBookParams, reply *repository.Book) error {
	book, err := s.Repo.UpdateBook(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = book

	return nil
}
