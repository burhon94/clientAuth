package client

import (
	"context"
	"errors"
	"github.com/burhon94/clientAuth/pkg/dl"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (c *Client) SignIn(ctx context.Context, clientRequest SignIn) (err error) {
	if clientRequest.Login == "" {
		return ErrBadRequest
	}

	if clientRequest.Pass == "" {
		return ErrBadRequest
	}

	err = c.CheckPassWithLogin(ctx, clientRequest.Login, clientRequest.Pass)
	if err != nil {
		switch {
		case errors.Is(err, ErrInternal):
			return ErrInternal

		case errors.Is(err, ErrInvalidLogin):
			return ErrInvalidLogin

		case errors.Is(err, ErrInvalidPassword):
			return ErrInvalidPassword

		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	var name, surName, middleName, eMail, avatar, phone string

	err = c.pool.QueryRow(ctx, dl.SignIn, clientRequest.Login).Scan(&name, &surName, &middleName, &eMail, &avatar, &phone)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrInvalidLogin

		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx

		}

		return ErrInternal
	}

	return nil
}

func (c *Client) NewClient(ctx context.Context, clientData NewClientStruct) (err error) {
	if clientData.Login == "" {
		return ErrBadRequest
	}

	err = c.CheckLogin(ctx, clientData.Login)
	if err != nil {
		switch {
		case errors.Is(err, ErrLoginExist):
			return ErrLoginExist

		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx

		case errors.Is(err, ErrInternal):
			return ErrInternal
		}

		return ErrInternal
	}

	if clientData.Phone == "" {
		return ErrBadRequest
	}

	err = c.CheckPhone(ctx, clientData.Phone)
	if err != nil {
		switch {
		case errors.Is(err, ErrPhoneRegistered):
			return ErrPhoneRegistered

		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx

		case errors.Is(err, ErrInternal):
			return ErrInternal
		}

		return ErrInternal
	}

	if clientData.FirstName == "" {
		return ErrBadRequest
	}

	if clientData.LastName == "" {
		return ErrBadRequest
	}

	if clientData.MiddleName == "" {
		clientData.MiddleName = "-"
	}

	if clientData.Password == "" {
		return ErrBadRequest
	}

	password, err := bcrypt.GenerateFromPassword([]byte(clientData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("can't hashing password: %v", err)
		return err
	}

	if clientData.EMail == "" {
		clientData.EMail = "-"
	}

	if clientData.Avatar == "" {
		clientData.Avatar = "NO-AVATAR"
	}

	_, err = c.pool.Exec(ctx, dl.ClientNew,
		clientData.FirstName, clientData.LastName, clientData.MiddleName, clientData.Login, password, clientData.EMail, clientData.Avatar, clientData.Phone)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	return nil
}

func (c *Client) EditClientPass(ctx context.Context, clientData EditClientPass) (err error) {
	if clientData.Id <= 0 {
		return ErrBadRequest
	}

	if clientData.OldPass == "" {
		return ErrBadRequest
	}

	if clientData.NewPass == "" {
		return ErrBadRequest
	}

	err = c.CheckId(ctx, clientData.Id)
	if err != nil {
		switch {
		case errors.Is(err, ErrInternal):
			return ErrInternal

		case errors.Is(err, errId):
			return ErrBadRequest

		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	oldPass, err := c.CheckPass(ctx, clientData.Id)
	if err != nil {
		switch {
		case errors.Is(err, errId):
			return ErrBadRequest

		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx

		case errors.Is(err, ErrInternal):
			return ErrInternal
		}

		return ErrInternal
	}

	err = bcrypt.CompareHashAndPassword([]byte(oldPass), []byte(clientData.OldPass))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ErrInvalidPassword
	}

	password, err := bcrypt.GenerateFromPassword([]byte(clientData.NewPass), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("can't hashing password: %v", err)
		return ErrInternal
	}

	_, err = c.pool.Exec(ctx, dl.ClientUpdatePass, clientData.Id, password)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	return nil
}

func (c *Client) EditClientAvatar(ctx context.Context, clientId int64, avatarUrl string) (err error) {
	if clientId <= 0 {
		return ErrBadRequest
	}

	if avatarUrl == "" {
		return ErrBadRequest
	}

	err = c.CheckId(ctx, clientId)
	if err != nil {
		switch {
		case errors.Is(err, ErrInternal):
			return ErrInternal

		case errors.Is(err, errId):
			return ErrBadRequest

		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	_, err = c.pool.Exec(ctx, dl.ClientUpdateAvatar, clientId, avatarUrl)
	if err != nil {
		switch {
		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	return nil
}
