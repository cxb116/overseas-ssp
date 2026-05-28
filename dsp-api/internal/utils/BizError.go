package utils

type BizError struct {
	Code int
	Msg  string
}

func (err BizError) Error() string {
	return err.Msg
}
