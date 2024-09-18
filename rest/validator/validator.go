package validator

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	zh_tw "github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
	"github.com/pengcainiao/zero/core/env"
	"github.com/pengcainiao/zero/core/logx"
	"github.com/pengcainiao/zero/core/stores/redis"
	"github.com/pengcainiao/zero/rest/httprouter"
	"github.com/pengcainiao/zero/tools"
	"github.com/pengcainiao/zero/tools/syncer"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

var (
	once  sync.Once
	valid *Validator
)

// Validator 定义验证器
type Validator struct {
	Validate *validator.Validate
	uni      *ut.UniversalTranslator
	locale   string
	trans    ut.Translator
	mp       *miniprogram.MiniProgram
}

// NewValidator 实例化验证器
func NewValidator() *Validator {
	once.Do(func() {
		var redisCache = redis.RedisCache{RedisNode: syncer.Redis()}
		// 初始化gin validator
		validate, _ := binding.Validator.Engine().(*validator.Validate)
		_ = validate.RegisterValidation("validTime", validTime)
		v := &Validator{
			Validate: validate,
			uni: ut.New(
				en.New(),
				zh.New(),
				zh_tw.New(),
			),
			mp: wechat.NewWechat().GetMiniProgram(&config.Config{
				AppID:     env.WechatAppID,
				AppSecret: env.WechatAppSecret,
				Cache:     &redisCache,
			}),
		}

		// 初始化默认语言
		v.SetLocale("zh")

		// 注册函数
		v.Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("json")
			if label == "" {
				label = field.Tag.Get("form")
			}
			return label + " "
		})

		valid = v
	})
	return valid
}

// SetLocale 设置语言
func (v *Validator) SetLocale(locale string) {
	v.locale = locale
	v.trans, _ = v.uni.GetTranslator(v.locale)
	switch v.locale {
	case "en":
		_ = en_translations.RegisterDefaultTranslations(v.Validate, v.trans)
	case "zh_tw":
		_ = zh_tw_translations.RegisterDefaultTranslations(v.Validate, v.trans)
	default:
		_ = zh_translations.RegisterDefaultTranslations(v.Validate, v.trans)
	}
}

// Translate 错误转换
func (v *Validator) Translate(errs error) string {
	switch err := errs.(type) {
	case validator.ValidationErrors:
		return err[0].Translate(v.trans)
	default:
		return errs.Error()
	}
}

// VerifyContent 验证内容
func (v *Validator) VerifyContent(content []byte) bool {
	accessToken, err := v.mp.GetContext().GetAccessToken()
	if err != nil {
		logx.NewTraceLogger(context.Background()).Err(err).Str("access_token", accessToken).Msg("验证器获取access_token失败")
		return true
	}

	req := tools.SendHTTP(httprouter.NewContextData(context.Background(), nil), "POST", fmt.Sprintf("https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s", accessToken))
	_, _ = req.JSONBody(map[string]interface{}{
		"content": string(content),
	})

	var res struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	err = req.ToJSON(&res)
	if err != nil {
		str, _ := req.String()
		logx.NewTraceLogger(context.Background()).Err(err).Str("result", str).Msg("调用微信验证接口失败")
		return true
	}

	if res.ErrCode == 87014 {
		logx.NewTraceLogger(context.Background()).Err(errors.New(res.ErrMsg)).Interface("result", res).Msg("调用微信验证接口失败")
		return false
	}

	return true
}
