package events

import "github.com/pengcainiao2/zero/tools/sensors"

type WechatCuteGirlUse struct {
	sensors.BaseEvent
	AccepterID string `json:"accepter_id,omitempty"`
	SendType   string `json:"send_type,omitempty"`
}

func (w WechatCuteGirlUse) Name() string {
	return "wechat_cute_girl_use"
}

func (w WechatCuteGirlUse) Properties() map[string]interface{} {
	return sensors.Struct2Map(w)
}
