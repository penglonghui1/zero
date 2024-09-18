package events

import "github.com/pengcainiao2/zero/tools/sensors"

type AttentionOfficialAccountEvent struct {
	sensors.BaseEvent
	OfficialAccountName string `json:"official_account_name"`
	Origin              string `json:"origin"`
}

func (e AttentionOfficialAccountEvent) Name() string {
	return "subscribe_official_accounts_success"
}

func (e AttentionOfficialAccountEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
