package token

import (
	"errors"
	"time"
)
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)
type Maker interface{
	CreateToken(username string,duration time.Duration) (string,*Payload,error)
	VerifyToken(token string)(*Payload,error)
}