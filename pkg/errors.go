package pkg

import "errors"

var (
	//ErrInvalidEmailAndPass depicts wrong email and password combination
	ErrInvalidEmailAndPass = errors.New("error: Invalid email and password combination")
	//ErrWrongFormat depicts wrong json format sent
	ErrWrongFormat = errors.New("error: Invalid format of data sent")
)
