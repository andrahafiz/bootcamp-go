package transactions

type CreateTransactionsRequest struct {
	EmployeeId int `json:"employee_id"`
	MenuId     int `json:"menu_id"`
	Quantity   int `json:"quantity"`
}
