package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/burhon94/clientAuth/pkg/dl"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var ErrBadRequest = errors.New("bad request")
var ErrLoginExist = errors.New("login is exist")
var ErrPhoneRegistered = errors.New(fmt.Sprintf("%s", "ERROR: duplicate key value violates unique constraint \"clients_phone_key\" (SQLSTATE 23505)"))

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
