package events

import "github.com/pengcainiao/zero/tools/sensors"

type LoginSuccessEvent struct {
	sensors.BaseEvent
	IsSuccess          bool   `json:"is_success"`            // 是否成功
	LoginWay           string `json:"login_way"`             // 登录类型
	IsClickLoginButton bool   `json:"is_click_login_button"` // 是否点击登录按钮
	FailReason         string `json:"fail_reason"`           // 失败原因
	IsNewUser          bool   `json:"is_new_user"`           // 是否新用户
}

func (event LoginSuccessEvent) Name() string {
	return "login_success"
}

func (event LoginSuccessEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(event)
}

func NewLoginEvent(isClickLoginButton, isSuccess bool, failReason string) LoginSuccessEvent {
	return LoginSuccessEvent{
		FailReason:         failReason,
		IsSuccess:          isSuccess,
		IsClickLoginButton: isClickLoginButton,
	}
}
