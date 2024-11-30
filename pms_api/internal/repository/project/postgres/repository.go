package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(p *pgxpool.Pool) *repository {
	return &repository{
		pool: p,
	}
}
