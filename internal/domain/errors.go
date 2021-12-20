package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("user doesn't exists")
	ErrUserAlreadyExists = errors.New("user with such email already exists")

	ErrCategoryNotFound      = errors.New("category doesn't exists")
	ErrCategoryAlreadyExists = errors.New("category with such title already exists")

	ErrListNotFound      = errors.New("list doesn't exists")
	ErrListAlreadyExists = errors.New("list with such title already exists")

	ErrItemNotFound      = errors.New("todo doesn't exists")
	ErrItemAlreadyExists = errors.New("todo with such title already exists")

	ErrVerificationCodeInvalid = errors.New("verification code is invalid")
)
