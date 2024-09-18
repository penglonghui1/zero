package events

import "github.com/pengcainiao2/zero/tools/sensors"

type PcLoginRoute struct {
	sensors.BaseEvent
	TwocodeID string `json:"twocode_id,omitempty"` // 二维码id
	Operate   string `json:"operate,omitempty"`    // 操作行为
}

func (event PcLoginRoute) Name() string {
	return "pc_login_route"
}

func (event PcLoginRoute) Properties() map[string]interface{} {
	return sensors.Struct2Map(event)
}
