package client

import (
	"context"
	"errors"
	"github.com/burhon94/clientAuth/pkg/dl"
	"github.com/burhon94/clientAuth/pkg/jwt"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
	"time"
)

func (c *Client) GenerateToken(ctx context.Context, clientRequest SignIn) (token Token, err error) {
	if clientRequest.Login == "" {
		return token, ErrBadRequest
	}

	if clientRequest.Pass == "" {
		return token, ErrBadRequest
	}

	err = c.CheckPassWithLogin(ctx, clientRequest.Login, clientRequest.Pass)
	if err != nil {
		switch {
		case errors.Is(err, ErrInternal):
			return token, ErrInternal

		case errors.Is(err, ErrInvalidLogin):
			return token, ErrInvalidLogin

		case errors.Is(err, ErrInvalidPassword):
			return token, ErrInvalidPassword

		case errors.Is(err, ErrTimeCtx):
			return token, ErrTimeCtx
		}

		return token, ErrInternal
	}

	var id int64

	err = c.pool.QueryRow(ctx, dl.Token, clientRequest.Login).Scan(&id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return token, ErrInvalidLogin

		case errors.Is(err, context.DeadlineExceeded):
			return token, ErrTimeCtx

		}

		return token, ErrInternal
	}

	token.Token, err = jwt.Encode(TokenPayload{
		Id:  id,
		Exp: time.Now().Add(time.Hour).Unix(),
	}, c.secret)
	if err != nil {
		return Token{}, ErrInternal
	}

	return
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

func (c *Client) EditClientPass(ctx context.Context, request *http.Request, clientData EditClientPass) (err error) {
	token, err := getToken(request)
	if err != nil {
		return ErrBadRequest
	}

	var tokenData TokenPayload
	err = jwt.Decode(token, &tokenData)
	if err != nil {
		return ErrBadRequest
	}

	if tokenData.Id <= 0 {
		return ErrBadRequest
	}

	if clientData.OldPass == "" {
		return ErrBadRequest
	}

	if clientData.NewPass == "" {
		return ErrBadRequest
	}

	err = c.CheckId(ctx, tokenData.Id)
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

	oldPass, err := c.CheckPass(ctx, tokenData.Id)
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

	_, err = c.pool.Exec(ctx, dl.ClientUpdatePass, tokenData.Id, password)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	return nil
}

func (c *Client) EditClientAvatar(ctx context.Context, request *http.Request, avatarUrl string) (err error) {
	token, err := getToken(request)
	if err != nil {
		return ErrBadRequest
	}

	var tokenData TokenPayload
	err = jwt.Decode(token, &tokenData)
	if err != nil {
		return ErrBadRequest
	}

	if tokenData.Id <= 0 {
		return ErrBadRequest
	}

	if avatarUrl == "" {
		return ErrBadRequest
	}

	err = c.CheckId(ctx, tokenData.Id)
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

	_, err = c.pool.Exec(ctx, dl.ClientUpdateAvatar, tokenData.Id, avatarUrl)
	if err != nil {
		switch {
		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	return nil
}

func (c *Client) EditClient(ctx context.Context, request *http.Request, requestData EditClient) error {
	token, err := getToken(request)
	if err != nil {
		return ErrBadRequest
	}

	var tokenData TokenPayload
	err = jwt.Decode(token, &tokenData)
	if err != nil {
		return ErrBadRequest
	}

	if tokenData.Id <= 0 {
		return ErrBadRequest
	}

	err = c.CheckId(ctx, tokenData.Id)
	if err != nil {
		switch {
		case errors.Is(err, errId):
			return ErrBadRequest

		case errors.Is(err, ErrInternal):
			return ErrInternal

		case errors.Is(err, ErrTimeCtx):
			return ErrTimeCtx
		}
	}

	if requestData.FirstName == "" {
		return ErrBadRequest
	}

	if requestData.LastName == "" {
		return ErrBadRequest
	}

	if requestData.MiddleName == "" {
		requestData.MiddleName = ""
	}

	if requestData.EMail == "" {
		requestData.EMail = ""
	}

	_, err = c.pool.Exec(ctx, dl.ClientUpdateData,
		tokenData.Id, requestData.FirstName, requestData.LastName, requestData.MiddleName, requestData.EMail)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return ErrTimeCtx
		}

		return ErrInternal
	}

	return nil
}

func getToken(request *http.Request) (token string, err error) {
	header := request.Header.Get("Authorization")
	if header == "" {
		return
	}

	if !strings.HasPrefix(header, "Bearer ") {
		return
	}

	token = header[len("Bearer "):]
	return
}
