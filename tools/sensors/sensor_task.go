package sensors

import (
	"context"

	"github.com/pengcainiao/zero/core/logx"
)

type EventType string

const (
	SignUp            EventType = "su"
	Normal            EventType = "n"
	UpdateUserProfile EventType = "up"
	ItemSet           EventType = "is"
)

type SensorTask struct {
	IdentityID  string
	AnonymousID string
	Event       Event
	Type        EventType
	EventName   string
}

func enQueue(task SensorTask) {
	_ = StartPushSensorsData(task)
}

// TrackEvent 神策事件上报
func (sensor *Sensor) TrackEvent(task SensorTask, isLogin bool) error {
	if sensor == nil {
		logx.NewTraceLogger(context.Background()).Warn().Msg("sensor SDK尚未初始化")
		return nil
	}
	event := task.Event
	event.FillLib(sensor.config.Lib)
	err := sensor.analytics.Track(task.IdentityID, event.Name(), event.Properties(), isLogin)

	if err != nil {
		trackFailed(task, err)
	}
	return err
}
