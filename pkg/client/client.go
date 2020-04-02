package client

import (
	"context"
	"fmt"
	"github.com/burhon94/clientAuth/pkg/dl"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
)

type Client struct {
	pool *pgxpool.Pool
}

func NewClient(pool *pgxpool.Pool) *Client {
	return &Client{pool: pool}
}

func (c *Client) Start()  {
	_, err := c.pool.Exec(context.Background(), dl.ClientDDL)
	if err != nil {
		panic(fmt.Sprintf("can't init DB: %v", err))
	}

	_, err = c.pool.Exec(context.Background(), dl.ClientDML)
	if err != nil {
		panic(fmt.Sprintf("can't set DB: %v", err))
	}
}

func (c *Client) NewClient(ctx context.Context, name, lastName, login, pass string, request *http.Request) (err error) {
	_, err = c.pool.Exec(context.Background(), dl.ClientNew, name, lastName, login, pass)
	if err != nil {
		log.Printf("can't insert new client: %v", err)
		return err
	}

	return nil
}