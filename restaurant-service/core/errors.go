package core

import "errors"

var (
	ErrInvalidRequestData        = errors.New(INVALID_REQUEST_DATA)
	ErrUnableToReadAllMenuFromDb = errors.New(UNABLE_TO_READ_ALL_MENU_FROM_DB)
	ErrUnableToAddMenuToDb       = errors.New(UNABLE_TO_ADD_MENU_TO_DB)
)
