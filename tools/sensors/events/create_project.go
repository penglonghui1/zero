package events

import "github.com/pengcainiao/zero/tools/sensors"

type CreateProjectEvent struct {
	sensors.BaseEvent
	ProjectParticipantQuantity int `json:"project_participant_quantity"` // 项目参与者数量
	TitleCharactersNumber      int `json:"title_characters_number"`      // 标题字符数量
	DetailsCharactersNumber    int `json:"details_characters_number"`    // 项目详情字符数量
	ProjectIssueNumber         int `json:"project_issue_number"`         // 项目事项总数
}

func (e CreateProjectEvent) Name() string {
	return "create_project"
}

func (e CreateProjectEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
