package domainerrors

import "errors"

var ErrDatabaseSave error = errors.New("error attempting to save on database")
var ErrInternal error = errors.New("Sorry we have experienced an unexpected error 😞")
