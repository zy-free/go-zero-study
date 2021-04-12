package ecode

// All common ecode
var (
	OK                 = add(0, "success")           	// 正确
	RequestErr         = add(400, "请求错误")        	// 请求错误
	Unauthorized       = add(401, "未认证")          	// 未认证
	AccessDenied       = add(403, "访问权限不足")     	// 访问权限不足
	NothingFound       = add(404, "啥都木有")        	// 啥都木有
	MethodNotAllowed   = add(405, "不支持该方法")      	// 不支持该方法
	ServerErr          = add(500, "服务器错误")       	// 服务器错误
	ServiceUnavailable = add(503, "过载保护,服务暂不可用") // 过载保护,服务暂不可用
	Deadline           = add(504, "服务调用超时")      	// 服务调用超时
	LimitExceed        = add(509, "超出限制")        	// 超出限制
)
