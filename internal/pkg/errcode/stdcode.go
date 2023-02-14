package errcode

var (
	OK                  = NewErrcode(200, "0", "ok")      // OK 成功
	InternalServerError = NewErrcode(200, "1", "服务器内部错误") // InternalServerError 服务器错误

)
