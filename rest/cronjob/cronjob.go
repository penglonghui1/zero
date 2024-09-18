package cronjob

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/pengcainiao/zero/core/logx"
	"github.com/pengcainiao/zero/tools"

	"github.com/pengcainiao/zero/core/env"
	"github.com/pengcainiao/zero/rest/httprouter"
	"github.com/pengcainiao/zero/rest/httpx"
)

var (
	// Month 月份
	Month = map[int]string{
		1:  "jan",
		2:  "feb",
		3:  "mar",
		4:  "apr",
		5:  "may",
		6:  "jun",
		7:  "jul",
		8:  "aug",
		9:  "sep",
		10: "oct",
		11: "nov",
		12: "dec",
	}

	// Weekday 周
	Weekday = map[int]string{
		0: "sun",
		1: "mon",
		2: "tue",
		3: "wed",
		4: "thu",
		5: "fri",
		6: "sat",
	}
)

// Cronjob 定时任务
type Cronjob struct {
	Code            int         `json:"code"`                 // code码
	RefID           string      `json:"ref_id"`               // 事项id
	RepeatType      int         `json:"-"`                    // 重复类型
	SchedulePrefix  string      `json:"-"`                    // 表达式前缀
	Schedule        interface{} `json:"schedule,omitempty"`   // 执行时间 cron 表达式
	MustAfter       int64       `json:"must_after,omitempty"` // 开始执行时间
	Expr            []string    `json:"expr,omitempty"`       // 执行时间 cron 表达式（多个）
	ExecuteMetadata interface{} `json:"metadata"`             // 执行数据
	Executor        string      `json:"executor"`             // 执行器   http, message-pub

	// 下面的后面废弃掉
	Name       string `json:"-"` // 任务名称
	Pid        string `json:"-"` // pid
	RetryTimes int    `json:"-"` // 重试次数
	Remark     string `json:"-"` // 备注
}

// NewCronjob 实例化
func NewCronjob() *Cronjob {
	return &Cronjob{
		RetryTimes: 3,
	}
}

// SetName 设置名称
func (c *Cronjob) SetName(name string) *Cronjob {
	c.Name = name
	return c
}

// SetRepeatType 设置重复类型
func (c *Cronjob) SetRepeatType(repeatType int) *Cronjob {
	c.RepeatType = repeatType
	return c
}

// SetSchedule 设置执行时间
func (c *Cronjob) SetSchedule(schedule interface{}) *Cronjob {

	var t time.Time

	switch schedule := schedule.(type) {
	case int:
		t = time.Unix(int64(schedule), 0)
	case int64:
		t = time.Unix(schedule, 0)
	case string:
		if schedule == "@now" {
			c.Schedule = schedule
		} else {
			t, _ = time.Parse("2006-01-02 15:04:05", schedule)
		}
	}

	if !t.IsZero() {
		c.MustAfter = tools.Strtotime(t.Format("2006-01-02"))
		switch c.RepeatType {
		case 1, 999:
			// 每天重复
			c.Schedule = fmt.Sprintf("%d %d %d ? ? ?", t.Second(), t.Minute(), t.Hour())
		case 2:
			// 每周重复
			c.Schedule = fmt.Sprintf("%d %d %d ? ? %s", t.Second(), t.Minute(), t.Hour(), Weekday[int(t.Weekday())])
		case 3:
			// 每两周重复
			c.Schedule = fmt.Sprintf("@every-2-weeks %d %d %d ? ? %s", t.Second(), t.Minute(), t.Hour(), Weekday[int(t.Weekday())])
		case 4:
			// 工作日重复
			c.Schedule = fmt.Sprintf("@weekday %d %d %d ? ? ?", t.Second(), t.Minute(), t.Hour())
		case 5:
			// 非工作日重复
			c.Schedule = fmt.Sprintf("@none-weekday %d %d %d ? ? ?", t.Second(), t.Minute(), t.Hour())
		case 6:
			// 每月重复
			c.Schedule = fmt.Sprintf("%d %d %d %d ? ?", t.Second(), t.Minute(), t.Hour(), t.Day())
		default:
			c.Schedule = fmt.Sprintf("@once %d %d %d %d %s ?", t.Second(), t.Minute(), t.Hour(), t.Day(), Month[int(t.Month())])
		}
	}

	return c
}

// SetExpr 设置多个表达式
func (c *Cronjob) SetExpr(exprs interface{}) *Cronjob {
	switch exprs := exprs.(type) {
	case []int:
		for _, schedule := range exprs {
			c.SetSchedule(schedule)
			if c.Schedule != nil {
				c.Expr = append(c.Expr, c.Schedule.(string))
			}
			c.Schedule = nil
		}
	case []int64:
		for _, schedule := range exprs {
			c.SetSchedule(schedule)
			if c.Schedule != nil {
				c.Expr = append(c.Expr, c.Schedule.(string))
			}
			c.Schedule = nil
		}
	}
	return c
}

// SetExecuteMetadata 设置执行数据
func (c *Cronjob) SetExecuteMetadata(executeMetadata interface{}) *Cronjob {
	switch executeMetadata.(type) {
	case *HTTPMetadata:
		c.Executor = "http"
	case *MessagePubMetadata:
		c.Executor = "message-pub"
	}
	c.ExecuteMetadata = executeMetadata
	return c
}

// SetRetryTimes 设置重试次数
func (c *Cronjob) SetRetryTimes(retryTimes int) *Cronjob {
	c.RetryTimes = retryTimes
	return c
}

// SetPid 设置PID
func (c *Cronjob) SetPid(pid string) *Cronjob {
	c.Pid = pid
	return c
}

// SetRemark 设置备注
func (c *Cronjob) SetRemark(remark string) *Cronjob {
	c.Remark = remark
	return c
}

// GetCronjobData 获取定时数据
func (c *Cronjob) GetCronjobData(ctx *httprouter.Context) Cronjob {
	if message, ok := c.ExecuteMetadata.(*MessagePubMetadata); ok {
		// message 设置用户信息
		c.Code = message.Code
		c.RefID = message.RefID
	} else if http, ok := c.ExecuteMetadata.(*HTTPMetadata); ok {
		// http 设置用户信息
		c.Code = http.Code
		c.RefID = http.RefID
		http.SetHeader("X-Auth-User", ctx.Data.UserID)
		http.SetHeader("X-Auth-Platform", ctx.Data.Platform)
		http.SetHeader("X-Auth-Version", ctx.Data.ClientVersion)
	}
	var cronjob Cronjob
	_ = tools.MapToStruct(c, &cronjob)
	return cronjob
}

// Send 推送任务
func (c *Cronjob) Send(ctx *httprouter.Context) error {
	cronjobURL := fmt.Sprintf("http://timedtask-svc.%s.svc.cluster.local:8080/v1/jobs", env.Namespace)
	if host := os.Getenv("TIMEDTASK_HOST"); host != "" {
		cronjobURL = host + "/jobs"
	}
	req := httpx.SendHTTP(ctx, "POST", cronjobURL)

	if message, ok := c.ExecuteMetadata.(*MessagePubMetadata); ok {
		// message 设置用户信息
		c.Code = message.Code
		c.RefID = message.RefID
	} else if http, ok := c.ExecuteMetadata.(*HTTPMetadata); ok {
		// http 设置用户信息
		c.Code = http.Code
		c.RefID = http.RefID
		http.SetHeader("X-Auth-User", ctx.Data.UserID)
		http.SetHeader("X-Auth-Platform", ctx.Data.Platform)
		http.SetHeader("X-Auth-Version", ctx.Data.ClientVersion)
	}

	_, _ = req.JSONBody(c)

	if os.Getenv("PRINT_SEND_CONTENT") == "true" {
		logx.NewTraceLogger(ctx).Debug().Interface("body", c).Msg("消息推送内容")
	}
	var res struct {
		Code    int         `json:"code"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}

	err := req.ToJSON(&res)
	if err != nil {
		logx.NewTraceLogger(ctx).Err(err).Interface("body", c).Msg("创建定时任务失败1")
		return err
	}

	if res.Code != 0 {
		logx.NewTraceLogger(ctx).Err(errors.New(res.Message)).Interface("body", c).Msg("创建定时任务失败2")
		return errors.New(res.Message)
	}
	return nil
}

// SendBatch 推送批量任务
func (c *Cronjob) SendBatch(ctx *httprouter.Context, cronjobs []Cronjob) error {
	cronjobURL := fmt.Sprintf("http://timedtask-svc.%s.svc.cluster.local:8080/v1/jobs/batch", env.Namespace)
	if host := os.Getenv("TIMEDTASK_HOST"); host != "" {
		cronjobURL = host + "/jobs/batch"
	}
	req := httpx.SendHTTP(ctx, "POST", cronjobURL)
	_, _ = req.JSONBody(cronjobs)

	if os.Getenv("PRINT_SEND_CONTENT") == "true" {
		logx.NewTraceLogger(ctx).Debug().Interface("body", c).Msg("消息推送内容")
	}
	var res struct {
		Code    int         `json:"code"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}

	err := req.ToJSON(&res)
	if err != nil {
		logx.NewTraceLogger(ctx).Err(err).Interface("body", c).Msg("创建定时任务失败1")
		return err
	}

	if res.Code != 0 {
		logx.NewTraceLogger(ctx).Err(errors.New(res.Message)).Interface("body", c).Msg("创建定时任务失败2")
		return errors.New(res.Message)
	}
	return nil
}
