package sensors

import (
	"context"

	"github.com/pengcainiao2/zero/core/logx"
	"github.com/pengcainiao2/zero/rest/httpx"
)

var (
	hook = httpx.WebHook()
)

// SensorsClient 神策客户端
func SensorsClient() *Sensor {
	return SensorDataProducer()
}

// StartPushSensorsData 开始推送神策数据，应包含与 nsq的HandleMessage方法内
func StartPushSensorsData(s SensorTask) error {
	var (
		err           error
		sensorsClient = SensorDataProducer()
	)
	switch s.Type {
	case Normal:
		if err = sensorsClient.TrackEvent(s, !s.Event.IsAnonymous()); err != nil {
			logx.NewTraceLogger(context.Background()).Err(err).
				Str("user_id", s.IdentityID).
				Msg("记录事件失败")
		}
	case SignUp:
		if err = sensorsClient.Analytics().TrackSignup(s.IdentityID, s.AnonymousID); err != nil {
			logx.NewTraceLogger(context.Background()).Err(err).
				Str("user_id", s.IdentityID).
				Str("anonymous_id", s.AnonymousID).
				Msg("记录账号关联失败")
		}
	case UpdateUserProfile:
		data := s.Event.Properties()
		if err = sensorsClient.Analytics().ProfileSet(s.IdentityID, data, true); err != nil {
			logx.NewTraceLogger(context.Background()).Err(err).
				Str("user_id", s.IdentityID).
				Str("anonymous_id", s.AnonymousID).
				Interface("data", data).
				Msg("记录账号关联失败")
		}
	case ItemSet:
		data := s.Event.Properties()
		if err = sensorsClient.Analytics().ItemSet(s.Event.Name(), s.IdentityID, data); err != nil {
			logx.NewTraceLogger(context.Background()).Err(err).
				Str("name", s.Event.Name()).
				Str("item_id", s.IdentityID).
				Interface("item_data", data).
				Msg("记录纬度表设置失败")
		}
	}
	return err
}

func trackFailed(task SensorTask, err error) {
	var reperr = httpx.ReportErrors{
		Title: "神策埋点错误",
		Error: err,
		Extra: map[string]interface{}{
			"event":         task.Event.Name(),
			"task_identity": task.IdentityID,
			"anonymous_id":  task.AnonymousID,
		},
	}
	logx.NewTraceLogger(context.Background()).Err(err).Str("event", task.Event.Name()).
		Str("user_id", task.IdentityID).
		Str("anonymous_id", task.AnonymousID).Msg("上报出错")
	hook.MarkdownReport(reperr)
}
