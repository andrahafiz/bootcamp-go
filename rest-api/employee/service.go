package employee

import (
	"log"
)

type RepositoryInterface interface {
	Save(employee Employee) (err error)
	GetAll() (employees []Employee, err error)
	GetById(id int) (employee Employee, err error)
	Delete(id int) (err error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Save(employee Employee) (err error) {
	err = s.repo.Save(employee)
	if err != nil {
		log.Println("error when try to save menu with error", err.Error())
		return
	}
	return
}
