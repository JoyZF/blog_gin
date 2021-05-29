package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(100000000, "服务内部错误")
	InvalidParams = NewError(100000001, "入参错误")
	NotFound      = NewError(100000002, "未找到")
	UnAuth        = NewError(100000003, "鉴权失败")
	TooManyReq    = NewError(100000004, "请求过多")
)
