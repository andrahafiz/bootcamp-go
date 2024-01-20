package transactions

import (
	"database/sql"
	"errors"
	"log"
)

type RepositoryInterface interface {
	Save(data Transactions) (err error)
	GetMenuById(id int) (menu Menu, err error)
	GetEmployeeById(id int) (emp Employee, err error)
	GetAll() (data []Transactions, err error)
	FindById(id int) (data Transactions, err error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Process(req Transactions) (err error) {
	_, err = s.repo.GetEmployeeById(req.EmployeeId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("menu not found")
		}
		return
	}
	menu, err := s.repo.GetMenuById(req.MenuId)
	if err != nil {
		log.Println("error when try to GetMenuById with menu_id:", req.MenuId, "and error", err.Error())
		if err == sql.ErrNoRows {
			err = errors.New("menu not found")
		}
		return
	}

	req.SetTotal(menu.Price)

	err = s.repo.Save(req)
	return
}
