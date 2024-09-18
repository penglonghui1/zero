package events

import "github.com/pengcainiao2/zero/tools/sensors"

type PaySuccessBox struct {
	sensors.BaseEvent
	SourceType string `json:"source_type,omitempty"`
	OrderID    string `json:"order_id,omitempty"`
}

func (p PaySuccessBox) Name() string {
	return "pay_success_box"
}

func (p PaySuccessBox) Properties() map[string]interface{} {
	return sensors.Struct2Map(p)
}
