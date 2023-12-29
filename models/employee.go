package models

type Employee struct {
	tableName struct{} `pg:"employee_table"`
	ID        string   `json:"id" pg:"id,pk"`
	Name      string   `json:"name,omitempty" pg:"name"`
	HP        string   `json:"hp,omitempty" pg:"hp"`
	Status    string   `json:"status,omitempty" pg:"status"`
}
