package events

import "github.com/pengcainiao/zero/tools/sensors"

type OpenMemberNotice struct {
	sensors.BaseEvent
	BuyMemberType string `json:"buy_member_type,omitempty"`
	OrderID       string `json:"order_id,omitempty"`
	PayType       string `json:"pay_type,omitempty"`
}

func (o OpenMemberNotice) Name() string {
	return "open_member_notice"
}

func (o OpenMemberNotice) Properties() map[string]interface{} {
	return sensors.Struct2Map(o)
}
