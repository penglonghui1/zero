package sensors

import (
	"context"
	"net/url"
	"sync"

	"github.com/pengcainiao2/zero/core/env"
	"github.com/pengcainiao2/zero/core/logx"
	"github.com/pengcainiao2/zero/core/sysx"
	sdk "github.com/sensorsdata/sa-sdk-go"
	"github.com/sensorsdata/sa-sdk-go/consumers"
)

var (
	sensor   *Sensor
	initOnce sync.Once
)

type Sensor struct {
	consumer  consumers.Consumer
	analytics *sdk.SensorsAnalytics
	config    *ClientConfig
}

type ClientConfig struct {
	Lib *Lib
}

// SensorDataProducer 向神策服务器发送数据
func SensorDataProducer() *Sensor {
	// 从神策分析配置页面中获取数据接收的 URL
	initOnce.Do(func() {
		saURL := env.SensorsAddr
		parse, err := url.Parse(saURL)
		if err != nil || saURL == "" {
			logx.NewTraceLogger(context.Background()).Err(err).Str("SA_URL", saURL).Msg("无法初始化sensorsdata SDK")
			return
		}
		// 初始化一个 Consumer，用于数据发送
		// DefaultConsumer 是同步发送数据，因此不要在任何线上的服务中使用此 Consumer
		consumer, err := sdk.InitDefaultConsumer(saURL, 5000)
		//consumer, err := sdk.InitBatchConsumer(saURL, 30, 5000)
		//consumer, err := sdk.InitConcurrentLoggingConsumer(os.Getenv("SA_LOG_PATH"), true)

		if err != nil {
			logx.NewTraceLogger(context.Background()).Debug().Err(err).Str("SA_URL", parse.String()).Msg("神策初始化失败")
			return
		}
		// 使用 Consumer 来构造 SensorsAnalytics 对象
		sa := sdk.InitSensorsAnalytics(consumer, parse.Query().Get("project"), false)

		sensor = &Sensor{
			consumer:  consumer,
			analytics: &sa,
			config: &ClientConfig{
				Lib: &Lib{
					AppVersion: sysx.AppVersion,
				},
			},
		}
	})
	return sensor
}

// Client 初始化客户端
func Client() *Sensor {
	sensor.config = &ClientConfig{
		Lib: &Lib{
			AppVersion: sysx.AppVersion,
		},
	}
	return sensor
}

// Consumer 获取默认消费者
func (sensor *Sensor) Consumer() consumers.Consumer {
	if sensor == nil {
		logx.NewTraceLogger(context.Background()).Warn().Msg("sensor SDK尚未初始化")
		return nil
	}
	return sensor.consumer
}

// Analytics 获取分析器
func (sensor *Sensor) Analytics() *sdk.SensorsAnalytics {
	if sensor == nil {
		logx.NewTraceLogger(context.Background()).Warn().Msg("sensor SDK尚未初始化")
		return &sdk.SensorsAnalytics{}
	}
	return sensor.analytics
}

// Track 神策普通事件上报
func (sensor *Sensor) Track(userID string, event Event) {
	if sensor == nil {
		logx.NewTraceLogger(context.Background()).Warn().Msg("sensor SDK尚未初始化")
		return
	}
	enQueue(SensorTask{
		IdentityID: userID,
		Event:      event,
		Type:       Normal,
		EventName:  event.Name(),
	})
}

// ProfileSet 更新用户信息
func (sensor *Sensor) ProfileSet(userID string, profile Event) {
	if sensor == nil {
		logx.NewTraceLogger(context.Background()).Warn().Msg("sensor SDK尚未初始化")
		return
	}
	enQueue(SensorTask{
		IdentityID: userID,
		Event:      profile,
		Type:       UpdateUserProfile,
		EventName:  profile.Name(),
	})
}

// ItemSet 更新用户信息
func (sensor *Sensor) ItemSet(item Item) {
	if sensor == nil {
		logx.NewTraceLogger(context.Background()).Warn().Msg("sensor SDK尚未初始化")
		return
	}
	enQueue(SensorTask{
		IdentityID: item.GetItemID(),
		Event:      item,
		Type:       ItemSet,
		EventName:  item.Name(),
	})
}

// TrackSignup 注册关联
func (sensor *Sensor) TrackSignup(userID, anonymousID string) {
	if sensor == nil {
		return
	}
	enQueue(SensorTask{
		IdentityID:  userID,
		AnonymousID: anonymousID,
		Event:       &UserPairEvent{},
		Type:        SignUp,
		EventName:   "track_signup",
	})
}
