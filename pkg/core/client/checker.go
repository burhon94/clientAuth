package client

import (
	"context"
	"errors"
	"github.com/burhon94/clientAuth/pkg/dl"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func (c *Client) CheckId(ctx context.Context, checkId int64) error {
	var id int64
	err := c.pool.QueryRow(ctx, dl.CheckId, checkId).Scan(&id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return errId

		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	return nil
}

func (c *Client) CheckLogin(ctx context.Context, login string) error {
	temp := ""
	err := c.pool.QueryRow(ctx, dl.CheckLogin, login).Scan(&temp)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil

		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx

		}

		return ErrInternal
	}

	return ErrLoginExist
}

func (c *Client) CheckPhone(ctx context.Context, phone string) error {
	temp := ""
	err := c.pool.QueryRow(ctx, dl.CheckPhone, phone).Scan(&temp)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil

		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx

		}

		return ErrInternal
	}

	return ErrPhoneRegistered
}

func (c *Client) CheckPass(ctx context.Context, id int64) (oldPass string, err error) {
	err = c.pool.QueryRow(ctx, dl.CheckPass, id).Scan(&oldPass)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return "", errId

		case errors.Is(err, context.DeadlineExceeded):
			return "", ErrTimeCtx
		}

		return "", ErrInternal
	}

	return oldPass, nil
}

func (c *Client) CheckPassWithLogin(ctx context.Context, login, requiredPass string) (err error) {
	var pass string
	err = c.pool.QueryRow(ctx, dl.CheckPassAndLogin, login).Scan(&pass)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrInvalidLogin

		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx

		}

		return ErrInternal
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(requiredPass))
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}
