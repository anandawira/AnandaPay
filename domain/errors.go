package domain

import "errors"

/*
Internal Server Error
Email already in use
Email not found
Incorrect email or password
Parameter validation failed
*/

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrEmailUsed           = errors.New("email already in use")
	ErrEmailNotFound       = errors.New("email not found")
	ErrWalletNotFound      = errors.New("wallet not found")
	ErrWrongEmailPass      = errors.New("incorrect email or password")
	ErrParameterValidation = errors.New("parameter validation failed")
	ErrInvalidToken        = errors.New("jwt authentication token invalid")
	ErrInsufficientBalance = errors.New("wallet balance insufficient")
	ErrSendToOwn           = errors.New("cannot send money to your own wallet")
)
