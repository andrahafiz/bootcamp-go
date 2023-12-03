package auth

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api-2/utility"
)

type RepositoryInterface interface {
	Create(auth Auth) (err error)
	GetByEmail(email string) (auth Auth, err error)
	GetById(id any) (auth Auth, err error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Create(auth Auth) (err error) {
	auth.Password, err = utility.Hash(auth.Password)
	if err != nil {
		log.Println("error when try to hash password with error", err.Error())
		return
	}

	err = s.repo.Create(auth)
	if err != nil {
		log.Println("error when try to create auth with error", err.Error())
		return
	}
	return
}

func (s Service) Login(req Auth) (auth Auth, err error) {
	auth, err = s.repo.GetByEmail(req.Email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err.Error())
		if err == sql.ErrNoRows {
			err = fmt.Errorf("username or password not found")
			return
		}
		return
	}

	err = utility.Verify(auth.Password, req.Password)
	if err != nil {
		log.Println("error when try to verify password with error", err.Error())
		err = fmt.Errorf("username or password not found")
		return
	}

	return
}

func (s Service) GetById(id any) (auth Auth, err error) {
	auth, err = s.repo.GetById(id)
	if err != nil {
		log.Println("error when try to get auth with error", err.Error())
		return
	}

	return
}
