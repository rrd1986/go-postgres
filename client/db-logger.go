package client

import (
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	qbytes, err := q.FormattedQuery()
	if err != nil {
		return err
	}
	fmt.Println(string(qbytes))
	return nil
}
