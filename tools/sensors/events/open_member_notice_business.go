package events

import "github.com/pengcainiao/zero/tools/sensors"

type OpenMemberNoticeBusiness struct {
	sensors.BaseEvent
	BuyMemberType string `json:"buy_member_type,omitempty"`
	OrderID       string `json:"order_id,omitempty"`
	BuyDuration   string `json:"buy_duration,omitempty"`
	BuySeat       string `json:"buy_seat,omitempty"`
}

func (o OpenMemberNoticeBusiness) Name() string {
	return "open_member_notice_business"
}

func (o OpenMemberNoticeBusiness) Properties() map[string]interface{} {
	return sensors.Struct2Map(o)
}
