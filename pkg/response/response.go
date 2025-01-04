package response

import "github.com/gin-gonic/gin"

// Response 是统一的响应格式
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

func Success(c *gin.Context, code int, message string, data interface{}) {
	if code == 0 {
		code = 1
	}
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func Fail(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func SuccessWith(c *gin.Context, message string, data interface{}) {
	Success(c, 0, message, data)
}

func SuccessWithData(c *gin.Context, data interface{}) {
	Success(c, 0, "success", data)
}

func FailWithMessage(c *gin.Context, message string) {
	Fail(c, 1, message, nil)
}
