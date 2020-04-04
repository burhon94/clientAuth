package client

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

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

func (c *Client) CheckPass(ctx context.Context, id int64) (oldPass string, err error) {
	err = c.pool.QueryRow(ctx, `SELECT password from clients WHERE id = $1`, id).Scan(&oldPass)
	if err != nil {
		return "", err
	}

	return oldPass, nil
}

func (c *Client) CheckPassWithLogin(ctx context.Context, login, requiredPass string) (err error) {
	var pass string
	err = c.pool.QueryRow(ctx, `SELECT password from clients WHERE login = $1`, login).Scan(&pass)
	if err != nil {
		return ErrInvalidLogin
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(requiredPass))
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}