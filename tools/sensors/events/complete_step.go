package events

import "github.com/pengcainiao/zero/tools/sensors"

type CompleteStep struct {
	sensors.BaseEvent
	SpaceID          string `json:"space_id,omitempty"`
	ProjectID        string `json:"project_id,omitempty"`
	BusinessID       string `json:"business_id,omitempty"`
	CompleteEntrance string `json:"complete_entrance,omitempty"`
}

func (c CompleteStep) Name() string {
	return "complete_step"
}

func (c CompleteStep) Properties() map[string]interface{} {
	return sensors.Struct2Map(c)
}
