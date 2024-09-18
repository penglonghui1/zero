package events

import (
	"github.com/pengcainiao/zero/tools/sensors"
)

type UserProfile struct {
	Gender                      string `json:"gender,omitempty"`                         //性别
	NickName                    string `json:"nick_name,omitempty"`                      //昵称
	Country                     string `json:"country,omitempty"`                        //国家
	City                        string `json:"city,omitempty"`                           //城市
	Province                    string `json:"province,omitempty"`                       //省份
	IsFromInvite                string `json:"is_from_invite,omitempty"`                 //是否来源于邀请
	UserOrigin                  string `json:"user_origin,omitempty"`                    //用户来源
	WechatOrgin                 string `json:"wechat_orgin,omitempty"`                   //小程序渠道
	ChannelOrigin               string `json:"channel_origin,omitempty"`                 //APP渠道
	IsScheduleRemindSet         string `json:"is_schedule_remind_set,omitempty"`         //是否设置日程提醒
	RemindTime                  string `json:"remind_time,omitempty"`                    //今日日程提醒中发送时间的选择
	RemindTrigger               string `json:"remind_trigger,omitempty"`                 //提醒条件
	FileSyncPolicy              string `json:"file_sync_policy,omitempty"`               //同步策略
	TotalSpace                  string `json:"total_space,omitempty"`                    //单位 MB，保留三位小数
	UsedSpace                   string `json:"used_space,omitempty"`                     //单位 MB，保留三位小数
	IsOfficialAccountSubscribed string `json:"is_official_account_subscribed,omitempty"` //是否已经关注公众号
	PhoneNumber                 string `json:"phone_number,omitempty"`                   // 手机号
	LastUseWechatTime           string `json:"last_use_wechat_time,omitempty"`           // 最后一次使用小程序时间
	LastUseAppTime              string `json:"last_use_app_time,omitempty"`              // 最后一次使用APP时间
	LastUsePcTime               string `json:"last_use_pc_time,omitempty"`               // 最后一次使用PC时间
	SignupTime                  string `json:"signup_time,omitempty"`                    // 注册时间
	SuccessOrigin               string `json:"success_origin,omitempty"`                 // 注册渠道
	UserOriginType              string `json:"user_origin_type,omitempty"`               // 用户属性
	SuccessPlatformType         string `json:"success_platform_type,omitempty"`          // 注册平台类型
}

func (u UserProfile) Name() string {
	return "user_profile"
}

func (u UserProfile) IsAnonymous() bool {
	return false
}

func (u UserProfile) FillLib(lib *sensors.Lib) {

}

func (u UserProfile) Properties() map[string]interface{} {
	return sensors.Struct2Map(u)
}
