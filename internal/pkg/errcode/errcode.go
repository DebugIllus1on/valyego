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

// Decode 解析错误
func Decode(err error) (int, string, string) {
	// 如果为空则返回成功
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	// 判断错误类型
	switch typ := err.(type) {
	case *Errcode: // 类型是 Errcode
		return typ.HTTP, typ.Code, typ.Message
	default:
	}

	// 否则返回服务器错误
	return InternalServerError.HTTP, InternalServerError.Code, InternalServerError.Message
}
