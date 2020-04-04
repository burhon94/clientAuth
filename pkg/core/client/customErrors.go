package client

import "errors"

var ErrBadRequest = errors.New("bad request")

var ErrInternal = errors.New("internal error")

var ErrLoginExist = errors.New("login is exist")

var ErrPhoneRegistered = errors.New("phone been registered")

var ErrInvalidPassword = errors.New("invalid password")

var ErrInvalidLogin = errors.New("invalid login")

var errId = errors.New("no exist id")

var ErrTimeCtx = errors.New("deadline context")
