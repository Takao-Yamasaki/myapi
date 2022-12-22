package apperrors

type MyAppError struct {
	// ErrCode型のErrCodeフィールド
	ErrCode
	Message string
	Err error `json:"-"`
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// errors.Is/errors.Asを使えるようにUnwrapを使う
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
