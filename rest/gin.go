package rest

import (
	"os"

	"github.com/pengcainiao/zero/core/conf"
	"github.com/pengcainiao/zero/core/sysx"

	"github.com/gin-gonic/gin"
	"github.com/pengcainiao/zero/core/env"
	"github.com/pengcainiao/zero/core/service"
	"github.com/pengcainiao/zero/rest/handler"
)

//NewGinServer 新建gin服务器
func NewGinServer(option ...conf.Option) *gin.Engine {
	router := gin.New()
	if err := service.SetupDefaultConf(option...); err != nil {
		return nil
	}

	router.Use(handler.WrapHttpHandler(
		handler.RecoverHandler,
		//handler.PrometheusHandler(),
	))

	if env.IsRunningInK8s() {
		//router.Use(handler.WrapHttpHandler(handler.TimeoutHandler(time.Second * 10)))
	}
	router.Use(handler.TracingHandler(sysx.SubSystem))
	router.Use(handler.Request())
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output:    os.Stdout,
		SkipPaths: []string{"/v1/user/verify"},
	}))
	return router
}
