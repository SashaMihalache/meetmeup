package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

type DBLogger struct {
}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}
