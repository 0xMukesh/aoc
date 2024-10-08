package errorcodes

type Errorcode struct {
	Message string
}

func (e Errorcode) Error() string {
	return e.Message
}

func NewErrorcode(msg string) Errorcode {
	return Errorcode{
		Message: msg,
	}
}

var (
	ErrUnimplemented = NewErrorcode("unimplemented")
)
