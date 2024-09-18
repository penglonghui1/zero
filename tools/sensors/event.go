package sensors

import (
	"github.com/pengcainiao/zero/core/mapping"
)

type Event interface {
	Name() string
	FillLib(lib *Lib)
	Properties() map[string]interface{}
	IsAnonymous() bool
}

func Struct2Map(s interface{}) map[string]interface{} {
	m, _ := mapping.StructToMap(s)
	return m
}

type BaseEvent struct {
	Lib           *Lib   `json:"$lib"`
	PlatformType  string `json:"platform_type"`
	ClientVersion string `json:"client_version"`
	Ip            string `json:"$ip"`
	IsAnonymousID bool   `json:"-"`
	ChannelSource string `json:"channel_source"`
}

type Lib struct {
	AppVersion string `json:"app_version"`
}

func (e *BaseEvent) FillLib(lib *Lib) {
	if e.Lib == nil {
		e.Lib = lib
	}
}

func (e *BaseEvent) IsAnonymous() bool {
	if e == nil {
		return false
	}
	return e.IsAnonymousID
}
