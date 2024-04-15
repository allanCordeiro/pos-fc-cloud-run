package domain

import "errors"

var (
	ErrInvalidZipCode  = errors.New("invalid zip code")
	ErrZipCodeNotFound = errors.New("can not find zipcode")
)
