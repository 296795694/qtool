// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusContinue                      = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols            = 101 // RFC 7231, 6.2.2
	StatusProcessing                    = 102 // RFC 2518, 10.1
	StatusEarlyHints                    = 103 // RFC 8297
	StatusOK                            = 200 // RFC 7231, 6.3.1
	StatusCreated                       = 201 // RFC 7231, 6.3.2
	StatusAccepted                      = 202 // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo          = 203 // RFC 7231, 6.3.4
	StatusNoContent                     = 204 // RFC 7231, 6.3.5
	StatusResetContent                  = 205 // RFC 7231, 6.3.6
	StatusPartialContent                = 206 // RFC 7233, 4.1
	StatusMultiStatus                   = 207 // RFC 4918, 11.1
	StatusAlreadyReported               = 208 // RFC 5842, 7.1
	StatusIMUsed                        = 226 // RFC 3229, 10.4.1
	StatusMultipleChoices               = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently              = 301 // RFC 7231, 6.4.2
	StatusFound                         = 302 // RFC 7231, 6.4.3
	StatusSeeOther                      = 303 // RFC 7231, 6.4.4
	StatusNotModified                   = 304 // RFC 7232, 4.1
	StatusUseProxy                      = 305 // RFC 7231, 6.4.5
	_                                   = 306 // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect             = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect             = 308 // RFC 7538, 3
	StatusBadRequest                    = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                  = 401 // RFC 7235, 3.1
	StatusPaymentRequired               = 402 // RFC 7231, 6.5.2
	StatusForbidden                     = 403 // RFC 7231, 6.5.3
	StatusNotFound                      = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed              = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                 = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired             = 407 // RFC 7235, 3.2
	StatusRequestTimeout                = 408 // RFC 7231, 6.5.7
	StatusConflict                      = 409 // RFC 7231, 6.5.8
	StatusGone                          = 410 // RFC 7231, 6.5.9
	StatusLengthRequired                = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed            = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge         = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong             = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType          = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable  = 416 // RFC 7233, 4.4
	StatusExpectationFailed             = 417 // RFC 7231, 6.5.14
	StatusTeapot                        = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest            = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity           = 422 // RFC 4918, 11.2
	StatusLocked                        = 423 // RFC 4918, 11.3
	StatusFailedDependency              = 424 // RFC 4918, 11.4
	StatusTooEarly                      = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired               = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired          = 428 // RFC 6585, 3
	StatusTooManyRequests               = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge   = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons    = 451 // RFC 7725, 3
	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

var statusText = map[int]string{
	StatusContinue:                      "继续",
	StatusSwitchingProtocols:            "交换协议",
	StatusProcessing:                    "处理",
	StatusEarlyHints:                    "早期暗示",
	StatusOK:                            "请求成功",
	StatusCreated:                       "创建",
	StatusAccepted:                      "认可",
	StatusNonAuthoritativeInfo:          "非授权信息",
	StatusNoContent:                     "没有内容",
	StatusResetContent:                  "重置内容",
	StatusPartialContent:                "部分内容",
	StatusMultiStatus:                   "多重状态",
	StatusAlreadyReported:               "已报告",
	StatusIMUsed:                        "已经被使用",
	StatusMultipleChoices:               "多项选择",
	StatusMovedPermanently:              "永久移除",
	StatusFound:                         "发现",
	StatusSeeOther:                      "查看其他",
	StatusNotModified:                   "未修改",
	StatusUseProxy:                      "使用代理",
	StatusTemporaryRedirect:             "临时重定向",
	StatusPermanentRedirect:             "永久重定向",
	StatusBadRequest:                    "错误的请求",
	StatusUnauthorized:                  "请先登录",
	StatusPaymentRequired:               "要求付款",
	StatusForbidden:                     "访问被禁止",
	StatusNotFound:                      "无效的请求",
	StatusMethodNotAllowed:              "不允许的方法",
	StatusNotAcceptable:                 "不可接受",
	StatusProxyAuthRequired:             "需要代理身份验证",
	StatusRequestTimeout:                "请求超时",
	StatusConflict:                      "冲突",
	StatusGone:                          "完成",
	StatusLengthRequired:                "所需长度",
	StatusPreconditionFailed:            "先决条件失败",
	StatusRequestEntityTooLarge:         "请求实体太大",
	StatusRequestURITooLong:             "请求URI太长",
	StatusUnsupportedMediaType:          "不支持的媒体类型",
	StatusRequestedRangeNotSatisfiable:  "请求范围不满足",
	StatusExpectationFailed:             "异常失败",
	StatusTeapot:                        "I'm a teapot",
	StatusMisdirectedRequest:            "错误的请求",
	StatusUnprocessableEntity:           "不可处理实体",
	StatusLocked:                        "锁定",
	StatusFailedDependency:              "失败的依赖关系",
	StatusTooEarly:                      "太早了",
	StatusUpgradeRequired:               "需要升级",
	StatusPreconditionRequired:          "要求先决条件",
	StatusTooManyRequests:               "请求太多",
	StatusRequestHeaderFieldsTooLarge:   "请求头字段太长",
	StatusUnavailableForLegalReasons:    "因法律原因无法使用",
	StatusInternalServerError:           "内部服务器错误",
	StatusNotImplemented:                "未实施",
	StatusBadGateway:                    "网关超时",
	StatusServiceUnavailable:            "服务不可用",
	StatusGatewayTimeout:                "网关超时",
	StatusHTTPVersionNotSupported:       "不支持的HTTP版本",
	StatusVariantAlsoNegotiates:         "变体也可以协商",
	StatusInsufficientStorage:           "存储不足",
	StatusLoopDetected:                  "检测到循环",
	StatusNotExtended:                   "未扩展",
	StatusNetworkAuthenticationRequired: "需要网络身份验证",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
