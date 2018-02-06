package httpconst

const (
	STATUS_CODE_TYPE_OK            = "ok"
	STATUS_CODE_TYPE_INVALID_REQ   = "invalid_request"
	STATUS_CODE_TYPE_NOT_FOUND     = "not_found"
	STATUS_CODE_TYPE_FORBIDDEN     = "forbidden"
	STATUS_CODE_TYPE_SERVER_ERROR  = "server_error"
	STATUS_CODE_TYPE_INVALID_TIME  = "invalid_time"
	STATUS_CODE_TYPE_ACCESS_DENIED = "access_denied"
)

var (
	STATUS_CODE_TYPE_DESC = map[string]int32{
		STATUS_CODE_TYPE_OK:            200, // 成功
		STATUS_CODE_TYPE_INVALID_REQ:   400, // 请求缺少某个必需参数，包含一个不支持的参数或参数值，或者格式不正确
		STATUS_CODE_TYPE_NOT_FOUND:     404, // 请求失败，请求所希望得到的资源未被在服务器上发现。在参数相同的情况下，不应该重复请求
		STATUS_CODE_TYPE_FORBIDDEN:     403, // 用户没有对当前动作的权限，引导重新身份验证并不能提供任何帮助，而且这个请求也不应该被重复提交
		STATUS_CODE_TYPE_SERVER_ERROR:  500, // 服务器出现异常情况 可稍等后重新尝试请求，但需有尝试上限，建议最多3次，如一直失败，则中断并告知用户
		STATUS_CODE_TYPE_INVALID_TIME:  400, // 客户端时间不正确，应请求服务器时间重新构造
		STATUS_CODE_TYPE_ACCESS_DENIED: 401, // AccessToken访问拒绝
	}
)
