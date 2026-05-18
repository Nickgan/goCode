package resp

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type Code int

const (
	RoleErrCode    Code = 1001
	NetworkErrCode Code = 1002
)

var codeMap = map[Code]string{
	RoleErrCode:    "权限错误",
	NetworkErrCode: "网络错误",
}

func init() {
	// 可能是一个
}

func response(c *gin.Context, r Response) {
	c.JSON(200, r)
}

func Ok(c *gin.Context, data any, msg string) {
	response(c, Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

func OkWithData(c *gin.Context, data any) {
	response(c, Response{
		Code: 0,
		Msg:  "成功",
		Data: data,
	})
}

func OkWithMsg(c *gin.Context, msg string) {
	response(c, Response{
		Code: 0,
		Msg:  msg,
		Data: map[string]any{},
	})
}

// 很少用
func Fail(c *gin.Context, code int, data any, msg string) {
	response(c, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 用的最多的
func FailWithMsg(c *gin.Context, msg string) {
	response(c, Response{
		Code: 7,
		Msg:  msg,
		Data: map[string]any{},
	})
}

// 自定义Code异常
func FailWithCode(c *gin.Context, code Code) {
	msg, ok := codeMap[code]
	if !ok {
		msg = "未知错误"
	}
	response(c, Response{
		Code: 7,
		Msg:  msg,
		Data: map[string]any{},
	})
}
