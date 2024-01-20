package employee

type CreateEmployeeRequest struct {
	Nip     string `json:"nip"`
	Name 	string `json:"name"`
	Address  string `json:"address"`
}