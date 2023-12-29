package main

import (
	"fmt"

	"github.com/rrd1986/go-postgres/client"
	"github.com/rrd1986/go-postgres/models"
)

func main() {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")
	dbClient, err := client.NewPostgresClient("postgres", "postgres", address, "postgres")
	if err != nil {
		panic(err)
	}

	empList := []models.Employee{
		{
			ID:     "1",
			Name:   "John Doe",
			HP:     "123-456-7890",
			Status: "Active",
		},
		{
			ID:     "2",
			Name:   "Jane Smith",
			HP:     "987-654-3210",
			Status: "Inactive",
		},
		// Add more Employee instances as needed
	}

	err = dbClient.CreateEmployeeTable()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = dbClient.InsertEmployeeRows(empList)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Define conditions for the WHERE clause
	conditions := map[string]interface{}{
		"id":   "1",
		"name": "John Doe",
	}

	employees, err := dbClient.SelectEmployeeRows(conditions)
	if err != nil {
		fmt.Println("Error selecting rows:", err)
		return
	}

	fmt.Println(employees)

	updateValues := map[string]interface{}{
		"status": "Inactive",
		"hp":     "444-444-4444",
	}

	err = dbClient.UpdateEmployeeRows(updateValues, conditions)
	if err != nil {
		fmt.Println("Error updating employee:", err)
		return
	}

	employees, err = dbClient.SelectEmployeeRows(conditions)
	if err != nil {
		fmt.Println("Error selecting rows:", err)
		return
	}

	fmt.Println(employees)

	err = dbClient.DeleteEmployeeRows(conditions)
	if err != nil {
		fmt.Println("Error deleting rows:", err)
		return
	}

	conditions = map[string]interface{}{}

	employees, err = dbClient.SelectEmployeeRows(conditions)
	if err != nil {
		fmt.Println("Error selecting rows:", err)
		return
	}

	fmt.Println(employees)

}
