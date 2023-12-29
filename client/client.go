package client

import (
	"fmt"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/rrd1986/go-postgres/models"
)

type PostgresClientType interface {
	ConnectionClose() error
	CreateEmployeeTable() error
	InsertEmployeeRows(employee []models.Employee) error
	SelectEmployeeRows(conditions map[string]interface{}) ([]models.Employee, error)
	UpdateEmployeeRows(updateValues map[string]interface{}, conditions map[string]interface{}) error
	DeleteEmployeeRows(conditions map[string]interface{}) error
}

type PostgresClient struct {
	db *pg.DB
}

func NewPostgresClient(user, password, address, database string) (PostgresClientType, error) {
	options := &pg.Options{
		User:     user,
		Password: password,
		Addr:     address,
		Database: database,
		PoolSize: 50,
	}
	db := pg.Connect(options)
	if db == nil {
		return nil, fmt.Errorf("cannot connect to postgres")
	}
	// Enable logging for debugging
	db.AddQueryHook(dbLogger{})

	return &PostgresClient{db: db}, nil
}

func (c *PostgresClient) CreateEmployeeTable() error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	return c.db.Model(&models.Employee{}).CreateTable(opts)
}

func (c *PostgresClient) InsertEmployeeRows(tests []models.Employee) error {
	_, err := c.db.Model(&tests).Insert()
	return err
}

func (c *PostgresClient) ConnectionClose() error {
	err := c.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *PostgresClient) SelectEmployeeRows(conditions map[string]interface{}) ([]models.Employee, error) {
	var employees []models.Employee

	query := c.db.Model(&employees)

	for key, value := range conditions {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (c *PostgresClient) UpdateEmployeeRows(updateValues map[string]interface{}, conditions map[string]interface{}) error {
	query := c.db.Model(&models.Employee{})

	for key, value := range updateValues {
		query = query.Set(fmt.Sprintf("%s = ?", key), value)
	}

	for key, value := range conditions {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	_, err := query.Update()

	return err
}

func (c *PostgresClient) DeleteEmployeeRows(conditions map[string]interface{}) error {
	query := c.db.Model(&models.Employee{})

	for key, value := range conditions {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	_, err := query.Delete()
	return err
}
