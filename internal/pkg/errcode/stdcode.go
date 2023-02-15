package errcode

var (
	OK                  = NewErrcode(200, "100100", "ok")      // OK 成功
	InternalServerError = NewErrcode(200, "100101", "服务器内部错误") // InternalServerError 服务器错误
	ErrResourceNotFound = NewErrcode(404, "100102", "对象不存在")
	ErrInvalidParams    = NewErrcode(400, "100103", "无效的请求")
	ErrCarNotFound      = NewErrcode(404, "100201", "目标车辆不存在")
)
