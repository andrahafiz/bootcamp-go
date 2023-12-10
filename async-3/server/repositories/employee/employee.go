package repositories

import "async-3/server/models"

type EmployeeRepository interface {
	Save(employee *models.Employee) error

	FindAll() (*[]models.Employee, error)

	FindByID(id int) (*models.Employee, error)

	UpdateByID(employee *models.Employee) error

	DeleteByID(id int) error
}
