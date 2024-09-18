package cronjob

import (
	json "github.com/json-iterator/go"
)

// HTTPMetadata 定义http数据
type HTTPMetadata struct {
	Code       int               `json:"code"`                  // 状态
	RefID      string            `json:"ref_id"`                // 关联id
	URL        string            `json:"url"`                   // 请求url
	Method     string            `json:"method"`                // 请求方法
	Header     map[string]string `json:"header,omitempty"`      // 请求头
	Body       string            `json:"body"`                  // 请求body
	Timeout    int               `json:"timeout,omitempty"`     // 超时时间
	ExceptCode []string          `json:"except_code,omitempty"` // 请求返回的 HTTP_STATUS_CODE  200成功
}

// NewHTTPData 实例化
func NewHTTPData() *HTTPMetadata {
	return &HTTPMetadata{
		Method: "POST",
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       `{}`,
		ExceptCode: []string{"200"},
	}
}

// SetCode 设置code
func (h *HTTPMetadata) SetCode(code int) *HTTPMetadata {
	h.Code = code
	return h
}

// SetURL 设置url
func (h *HTTPMetadata) SetURL(url string) *HTTPMetadata {
	h.URL = url
	return h
}

// SetMethod 设置method
func (h *HTTPMetadata) SetMethod(method string) *HTTPMetadata {
	h.Method = method
	return h
}

// SetHeader 设置header
func (h *HTTPMetadata) SetHeader(key, value string) *HTTPMetadata {
	h.Header[key] = value
	return h
}

// SetBody 设置body
func (h *HTTPMetadata) SetBody(body map[string]interface{}) *HTTPMetadata {
	bodyStr, _ := json.Marshal(body)
	h.Body = string(bodyStr)
	return h
}

// SetTimeout 设置超时时间
func (h *HTTPMetadata) SetTimeout(timeout int) *HTTPMetadata {
	h.Timeout = timeout
	return h
}

// SetExceptCode 设置成功返回错误码
func (h *HTTPMetadata) SetExceptCode(exceptCode []string) *HTTPMetadata {
	h.ExceptCode = exceptCode
	return h
}
