package events

import "github.com/pengcainiao2/zero/tools/sensors"

type LogOutEvent struct {
	sensors.BaseEvent
	ActivePassive string `json:"active_passive,omitempty"` // 主动被动操作
}

func (event LogOutEvent) Name() string {
	return "log_out"
}

func (event LogOutEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(event)
}
