package events

import "github.com/pengcainiao/zero/tools/sensors"

type ReceiveMemberNotice struct {
	sensors.BaseEvent
	ReceiveContent    string `json:"receive_content,omitempty"`
	ReceiveDays       string `json:"receive_days,omitempty"`
	ReceiveMemberType string `json:"receive_member_type,omitempty"`
}

func (r ReceiveMemberNotice) Name() string {
	return "receive_member_notice"
}

func (r ReceiveMemberNotice) Properties() map[string]interface{} {
	return sensors.Struct2Map(r)
}
