package client

import (
	"context"
	"errors"
	"github.com/burhon94/clientAuth/pkg/dl"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var ErrBadRequest = errors.New("bad request")
var ErrInternal = errors.New("internal error")
var ErrLoginExist = errors.New("login is exist")
var ErrPhoneRegistered = errors.New("phone been registered")
var ErrInvalidPassword = errors.New("invalid password")
var ErrInvalidLogin = errors.New("invalid login")

func (c *Client) SignIn(ctx context.Context, clientRequest SignIn) (err error) {
	if clientRequest.Login == "" {
		return ErrBadRequest
	}

	if clientRequest.Pass == "" {
		return ErrBadRequest
	}

	var name, surName, midleName, eMail, avata, phone string

	err = c.CheckPassWithLogin(ctx, clientRequest.Login, clientRequest.Pass)
	switch {
	case errors.Is(err, ErrInternal):
		return ErrInternal


	case errors.Is(err, ErrInvalidLogin):
		return ErrInvalidLogin

	case errors.Is(err, ErrInvalidPassword):
		return ErrInvalidPassword
	}

	err = c.pool.QueryRow(ctx, dl.SignIn, clientRequest.Login).Scan(&name, &surName, &midleName, &eMail, &avata, &phone)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrInvalidLogin
		}

		return ErrInternal
	}

	return nil
}

func (c *Client) NewClient(ctx context.Context, clientData NewClientStruct) (err error) {
	if clientData.FirstName == "" {
		return ErrBadRequest
	}

	if clientData.LastName == "" {
		return ErrBadRequest
	}

	if clientData.MiddleName == "" {
		clientData.MiddleName = "-"
	}

	if clientData.Login == "" {
		return ErrBadRequest
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

	if clientData.Phone == "" {
		return ErrBadRequest
	}

	if clientData.Avatar == "" {
		clientData.Avatar = "NO-AVATAR"
	}

	err = c.CheckLogin(ctx, clientData.Login)
	if err != nil {
		return ErrLoginExist
	}

	err = c.CheckPhone(ctx, clientData.Phone)
	if err != nil {
		return ErrPhoneRegistered
	}

	_, err = c.pool.Exec(ctx, dl.ClientNew,
		clientData.FirstName, clientData.LastName, clientData.MiddleName, clientData.Login, password, clientData.EMail, clientData.Avatar, clientData.Phone)
	if err != nil {
		return err
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

	oldPass, err := c.CheckPass(ctx, clientData.Id)
	if err != nil {
		return ErrBadRequest
	}

	err = bcrypt.CompareHashAndPassword([]byte(oldPass), []byte(clientData.OldPass))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ErrInvalidPassword
	}

	password, err := bcrypt.GenerateFromPassword([]byte(clientData.NewPass), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("can't hashing password: %v", err)
		return err
	}

	_, err = c.pool.Exec(ctx, dl.ClientUpdatePass, clientData.Id, password)
	if err != nil {
		return err
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
		}
	}

	_, err = c.pool.Exec(ctx, dl.ClientUpdateAvatar, clientId, avatarUrl)
	if err != nil {
		return ErrInternal
	}

	return nil
}