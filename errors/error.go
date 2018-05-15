package error

import (
	"fmt"
	"time"
)

const (
	OK             = 0
	PARAMS_ERROR   = 10000
	NOTEXIST_ERROR = 10001
	AUTH_ERROR     = 20000
	REPASS_ERROR   = 20001
	NETWORK_ERROR  = 30000
	DB_ERROR       = 40000
	SERVER_ERROR   = 50000
)

/**
	GLOBLE RESULT CODE MAPPING
**/
var CODE_MAPPING map[int]string

func init() {
	CODE_MAPPING = map[int]string{
		0:     "成功",
		10000: "参数错误",
		10001: "数据不存在",
		20000: "访问受限",
		20001: "重复密码错误",
		30000: "网络错误",
		40000: "数据库错误",
		50000: "服务端错误",
	}
}

type ApiError struct {
	Code    int
	Message string
	Time    time.Time
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%d-%s", e.Code, e.Error)
}

func NewError(code int) ApiError {
	return NewErrorWithMessage(code, "发生未知错误")
}

func NewErrorWithMessage(code int, err string) ApiError {
	if v, ok := CODE_MAPPING[code]; ok {
		return ApiError{
			Code:    code,
			Message: v,
			Time:    time.Now(),
		}
	} else {
		return ApiError{
			Code:    code,
			Message: err,
			Time:    time.Now(),
		}
	}
}
