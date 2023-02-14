package errcode

import "fmt"

type Errcode struct {
	HTTP    int
	Code    string
	Message string
}

func NewErrcode(HTTP int, Code string, Message string) *Errcode {
	return &Errcode{
		HTTP, Code, Message,
	}
}

func (errcode *Errcode) Error() string {
	return errcode.Message
}

func (errcode *Errcode) SetMessage(stdfmt string, args ...interface{}) {
	errcode.Message = fmt.Sprintf(stdfmt, args...)
}

func Decode(err error) (int, string, string) {
	return 0, "", ""
}
