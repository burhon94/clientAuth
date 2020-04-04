package client

import (
	"context"
	"fmt"
	"github.com/burhon94/clientAuth/pkg/dl"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	pool *pgxpool.Pool
}

func NewClient(pool *pgxpool.Pool) *Client {
	return &Client{pool: pool}
}

func (c *Client) Start() {
	_, err := c.pool.Exec(context.Background(), dl.ClientDDL)
	if err != nil {
		panic(fmt.Sprintf("can't init DB: %v", err))
	}

	_, err = c.pool.Exec(context.Background(), dl.ClientDML)
	if err != nil {
		panic(fmt.Sprintf("can't set DB: %v", err))
	}
}
