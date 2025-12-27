package rpcservices

import (
	"net/http"

	"github.com/SomeSuperCoder/sqlclearning/internal/repository"
)

type UsersService struct {
	Repo *repository.Queries
}

func (s *UsersService) Insert(r *http.Request, args *repository.InsertUserParams, reply *repository.InsertUserRow) error {
	user, err := s.Repo.InsertUser(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = user

	return nil
}

func (s *UsersService) UpdateProfile(r *http.Request, args *repository.UpdateProfileParams, reply *repository.UpdateProfileRow) error {
	profile, err := s.Repo.UpdateProfile(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = profile

	return nil
}
