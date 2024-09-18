package events

import "github.com/pengcainiao/zero/tools/sensors"

type CreateIssueEvent struct {
	sensors.BaseEvent
	WorkspaceID              string `json:"space_id"`                     // 空间id
	ProjectID                string `json:"project_id"`                   // 项目id
	CreateProjectPeopleID    string `json:"create_project_people_id"`     // 项目创建人id
	IssueUserID              string `json:"issue_user_id"`                // 事项创建人id
	IssueId                  string `json:"issue_id"`                     // 事项id
	UpperLevelIssueID        string `json:"upper_level_issue_id"`         // 上级事项id
	BusinessType             string `json:"business_type"`                // 业务类型
	IsAddLabel               bool   `json:"is_add_label"`                 // 是否添加标签
	IsRepetitiveIssue        bool   `json:"is_repetitive_issue"`          // 是否是循环事项
	LabelNumber              int    `json:"label_number"`                 // 标签数量
	IssueHierarchyType       string `json:"issue_hierarchy_type"`         // 事项层级类型
	CreateSonIssuePeopleType string `json:"create_son_issue_people_type"` // 创建子事项用户类型
	EnterCreateIssuePage     string `json:"enter_create_issue_page"`      // 进入创建事项页面
	CreateType               string `json:"create_type"`                  // 创建类型
	IssueTitle               string `json:"issue_title"`                  // 事项标题
	IsStartAlarmTime         bool   `json:"is_start_alarm_time"`          // 是否有开始时间
	HasDeadline              bool   `json:"has_deadline"`                 // 是否有截止时间
	IsAlarmDeadline          bool   `json:"is_alarm_deadline"`            // 是否有提醒时间
}

func (c CreateIssueEvent) Name() string {
	return "create_issue"
}

func (c CreateIssueEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(c)
}
