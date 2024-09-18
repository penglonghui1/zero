package events

import "github.com/pengcainiao/zero/tools/sensors"

type UserOrigin struct {
	sensors.BaseEvent
	UserOrigin string `json:"user_origin"`
	IsNewUser  bool   `json:"is_new_user"`
}

func (e UserOrigin) Name() string {
	return "user_origin"
}

func (e UserOrigin) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
