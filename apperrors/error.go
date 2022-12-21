package apperrors

type MyAppError struct {
	// ErrCode型のErrCodeフィールド
	ErrCode
	Message string
	Err error
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}

// errors.Is/errors.Asを使えるようにUnwrapを使う
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
