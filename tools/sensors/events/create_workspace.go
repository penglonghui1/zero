package events

import "github.com/pengcainiao2/zero/tools/sensors"

type CreateWorkspaceEvent struct {
	sensors.BaseEvent
	SpaceID                          string `json:"space_id,omitempty"`                             // 空间ID
	SpaceType                        string `json:"space_type,omitempty"`                           // 空间类型
	SpaceTitle                       string `json:"space_title,omitempty"`                          // 空间标题
	CreateSpacePeopleID              string `json:"create_space_people_id,omitempty"`               // 创建空间用户ID
	TradeType                        string `json:"trade_type,omitempty"`                           // 所处行业
	TeamSize                         string `json:"team_size,omitempty"`                            // 团队规模
	SpaceCreateType                  string `json:"space_create_type,omitempty"`                    // 空间创建方式
	NumberOfSpaceParticipants        int    `json:"number_of_space_participants,omitempty"`         // 空间内部成员数
	NumberOfSpaceOutsideParticipants int    `json:"number_of_space_outside_participants,omitempty"` // 空间外部成员数
	NumberOfSpaceItems               int    `json:"number_of_space_items,omitempty"`                // 空间项目数
}

func (e CreateWorkspaceEvent) Name() string {
	return "create_space"
}

func (e CreateWorkspaceEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
