package service

import (
	"log"

	"github.com/pengcainiao/zero/core/conf"
	"github.com/pengcainiao/zero/core/env"
	"github.com/pengcainiao/zero/core/load"
	"github.com/pengcainiao/zero/core/logx"
	"github.com/pengcainiao/zero/core/prometheus"
	"github.com/pengcainiao/zero/core/stat"
	"github.com/pengcainiao/zero/core/sysx"
	"github.com/pengcainiao/zero/core/trace"
)

const (
	// DevMode means development mode.
	DevMode = "develop"
	// TestMode means test mode.
	TestMode = "release"
	// RtMode means regression test mode.
	RtMode = "rt"
	// PreMode means pre-release mode.
	PreMode = "uat"
	// ProMode means production mode.
	ProMode = "prod"
)

// A ServiceConf is a service config.
type ServiceConf struct {
	Name       string
	Log        logx.LogConf
	Mode       string            `json:",default=prod,options=develop|release|rt|uat|prod"`
	MetricsUrl string            `json:",optional"`
	Prometheus prometheus.Config `json:",optional"`
	Telemetry  trace.Config      `json:",optional"`
}

func SetupDefaultConf(opts ...conf.Option) error {
	var config = ServiceConf{
		Name: sysx.SubSystem,
		Mode: env.ReleaseMode,
		Prometheus: prometheus.Config{
			Host: "0.0.0.0",
			Port: 5000,
			Path: "/metrics",
		},
	}
	opt := conf.ApplyConfig(opts...)
	config.Telemetry = trace.Config{
		Name:     sysx.SubSystem,
		Endpoint: opt.JaegerCollectorAddress, //"https://jaeger-collector.flyele.vip/api/traces", //"jaeger-operator-headless.observability.svc.cluster.local:14250", //grpc协议
		Sampler:  opt.JaegerSamplerRate,
		Batcher:  "jaeger",
	}

	return config.SetUp()
}

// MustSetUp sets up the service, exits on error.
func (sc ServiceConf) MustSetUp() {
	if err := sc.SetUp(); err != nil {
		log.Fatal(err)
	}
}

// SetUp sets up the service.
func (sc ServiceConf) SetUp() error {
	//if len(sc.Log.ServiceName) == 0 {
	//	sc.Log.ServiceName = sc.Name
	//}
	//if err := logx.SetUp(sc.Log); err != nil {
	//	return err
	//}
	sc.Mode = env.ReleaseMode
	sc.initMode()
	prometheus.StartAgent(sc.Prometheus)

	if len(sc.Telemetry.Name) == 0 {
		sc.Telemetry.Name = sc.Name
	}
	trace.StartAgent(sc.Telemetry)

	if len(sc.MetricsUrl) > 0 {
		stat.SetReportWriter(stat.NewRemoteWriter(sc.MetricsUrl))
	}

	return nil
}

func (sc ServiceConf) initMode() {
	switch sc.Mode {
	case DevMode, TestMode, RtMode, PreMode:
		load.Disable()
		stat.SetReporter(nil)
	}
}
