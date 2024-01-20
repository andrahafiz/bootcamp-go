package employee

import (
	"database/sql"
	"fmt"
)

/*
membuat struct repository
struct ini mempunyai dependency ke *sql.DB
yang mana merupakan sebuah object koneksi database
*/
type Repository struct {
	db *sql.DB
}

/*
function untuk inisiasi repository dengan memasukkan dependency
ke dalam parameter.
*/
func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Save(employee Employee) (err error) {
	query := `
		INSERT INTO employees (name, address, nip, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(employee.Name, employee.Address, employee.Nip, employee.CreatedAt, employee.UpdatedAt)
	return
}

func (r Repository) GetById(id int) (employee Employee, err error) {
	query := `
		SELECT 
			*
		FROM
			employees
		WHERE id = $1
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(
		&employee.Id, &employee.Name, &employee.Nip, &employee.Address, &employee.CreatedAt, &employee.UpdatedAt,
	)

	return
}

func (r Repository) GetAll() (menus []Employee, err error) {
	query := `
		SELECT 
		id, name, nip
		, address
		, created_at, updated_at
		FROM
			employees
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	fmt.Print(rows)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		menu := Employee{}
		err = rows.Scan(
			&menu.Id, &menu.Name, &menu.Nip,
			&menu.Address,
			&menu.CreatedAt, &menu.UpdatedAt,
		)
		if err != nil {
			return
		}
		menus = append(menus, menu)
	}

	return

}

func (r Repository) Delete(id int) (err error) {
	query := `
		DELETE FROM employees
		WHERE id = $1
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	stmt.Exec(id)
	return
}
