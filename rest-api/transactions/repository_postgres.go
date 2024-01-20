package transactions

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Save(data Transactions) (err error) {
	query := `
		INSERT INTO transactions (employee_id,menu_id,quantity,total,created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(data.EmployeeId, data.MenuId, data.Quantity, data.Total, data.CreatedAt, data.UpdatedAt)
	return
}

func (r Repository) GetEmployeeById(id int) (emp Employee, err error) {
	query := `SELECT id, name FROM employees WHERE id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&emp.Id, &emp.Name,
	)
	return
}

func (r Repository) GetMenuById(id int) (menu Menu, err error) {
	query := `SELECT id, name, price FROM menus where id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&menu.Id, &menu.Name, &menu.Price,
	)
	return
}

func (r Repository) GetAll() (data []Transactions, err error) {
	query := `
	SELECT 
	s.id, s.employee_id, e.id, e.name, s.menu_id , m.id, m.name, m.price , s.quantity , s.total , s.created_at , s.updated_at 
	FROM
		transactions s
	inner join employees e on e.id = s.employee_id 
	inner join menus m on m.id = s.menu_id 
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		datas := Transactions{}
		err = rows.Scan(&datas.Id,
			&datas.EmployeeId,
			&datas.Employee.Id,
			&datas.Employee.Name,
			&datas.MenuId,
			&datas.Menu.Id,
			&datas.Menu.Name,
			&datas.Menu.Price,
			&datas.Quantity,
			&datas.Total,
			&datas.CreatedAt,
			&datas.UpdatedAt,
		)
		if err != nil {
			return
		}
		data = append(data, datas)

	}

	return

}

func (r Repository) FindById(id int) (data Transactions, err error) {
	query := `SELECT 
				s.id, s.employee_id, e.id, e.name, s.menu_id , m.id, m.name, m.price , s.quantity , s.total , s.created_at , s.updated_at 
				FROM
					transactions s
				inner join employees e on e.id = s.employee_id 
				inner join menus m on m.id = s.menu_id 
				WHERE s.id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&data.Id, &data.EmployeeId, &data.MenuId, &data.Quantity, &data.Total, &data.CreatedAt, &data.UpdatedAt,
		&data.Menu.Id, &data.Menu.Name, &data.Menu.Price,
		&data.Employee.Id, &data.Employee.Name,
	)
	return
}
