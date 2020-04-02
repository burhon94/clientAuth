package client

import "context"

func (c *Client) CheckLogin(ctx context.Context, login string) error {
	temp := ""
	err := c.pool.QueryRow(ctx, `SELECT login from clients WHERE login = $1`, login).Scan(&temp)
	if err == nil {
		return ErrLoginExist
	}

	return nil
}

func (c *Client) CheckPhone(ctx context.Context, phone string) error {
	temp := ""
	err := c.pool.QueryRow(ctx, `SELECT phone from clients WHERE phone = $1`, phone).Scan(&temp)
	if err == nil {
		return ErrPhoneRegistered
	}

	return nil
}
