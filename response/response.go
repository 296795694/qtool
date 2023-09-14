package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
)

const (
	SUCCESS = "success"
	ERROR   = "error"
	UUID    = "uuid"
)

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	Code       int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Status     string      `json:"status"`
	StatusDesc string      `json:"statusDesc"` // 提示信息
	Uuid       string      `json:"uuid"`
	Data       interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
	Request    string      `json:"request"`
}

// Success 返回成功信息
func Success(c *gin.Context, data interface{}) {
	code := http.StatusOK
	msg := StatusText(code)
	Json(c, code, msg, data)
}

// Error 返回错误信息
func Error(c *gin.Context, code int, statusDesc string) {
	Json(c, code, statusDesc, make(map[string]interface{}))
}

// ErrorMes返回错误信息
func ErrorMsg(err error) string {
	env := gin.Mode()
	if env == gin.ReleaseMode {
		return "网络请求失败"
	} else {
		return err.Error()
	}
}

// Json 标准返回结果数据结构封装
func Json(c *gin.Context, code int, statusDesc string, data interface{}) {
	status := SUCCESS
	if code != 200 {
		status = ERROR
	}
	req := c.FullPath()
	reqArr := strings.Split(req, "/")
	if len(reqArr) == 4 {
		req = path.Join(reqArr[2], reqArr[3])
	}
	resData := JsonResponse{
		Code:       code,
		Status:     status,
		StatusDesc: statusDesc,
		Data:       data,
		Uuid:       c.GetString(UUID),
		Request:    req,
	}
	c.JSON(StatusOK, resData)
}

// PrintLog
// @Description: 控制打印Log
// @Auth syq
// @Date 2021-11-26 10:55:49
// @param a
func PrintLog(a ...interface{}) {
	//fmt.Println(a...)
}
