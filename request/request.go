package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/296795694/qtool/constant"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// GetJson 获取post参数
func GetJson(c *gin.Context) (map[string]interface{}, error) {
	jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	return data, err
}

// PostJson
// @Description: post json 请求
// @Auth syq
// @Date 2021-12-11 14:12:00
// @param url
// @param data
// @return string
// @return error
func PostJson(url string, data map[string]interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		err = errors.New("数据格式不正确")
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// GetPageSize
// @Description: 格式化单页数量
// @Auth syq
// @Date 2021-11-11 09:39:13
// @param pageSize
// @return int
func GetPageSize(pageSize int) int {
	if pageSize <= 0 {
		pageSize = constant.DefaultPageSize
	}
	return pageSize
}

// GetPageIndex
// @Description: 格式化页码
// @Auth syq
// @Date 2021-11-11 09:41:59
// @param pageIndex
// @return int
func GetPageIndex(pageIndex int) int {
	if pageIndex <= 0 {
		pageIndex = constant.DefaultPageIndex
	}
	return pageIndex
}
