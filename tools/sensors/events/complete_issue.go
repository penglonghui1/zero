package events

import "github.com/pengcainiao2/zero/tools/sensors"

type CompleteIssueEvent struct {
	sensors.BaseEvent
	BusinessID            string       `json:"business_id"`              // 事项id
	BusinessType          BusinessType `json:"business_type"`            // 业务类型
	CompleteIdentityType  UserIdentity `json:"complete_identity_type"`   // 身份类型
	OperatorID            string       `json:"operator_id"`              // 操作用户
	CompleteEntrance      string       `json:"complete_entrance"`        // 完成入口
	IssueExecuteMinutes   int          `json:"issue_execute_minutes"`    // 历时多少分钟
	IsNewUser             bool         `json:"is_new_user"`              // 是否新用户
	WorkspaceID           string       `json:"space_id"`                 // 空间id
	ProjectID             string       `json:"project_id"`               // 项目id
	ProjectTitle          string       `json:"project_title"`            // 项目标题
	IsRepetitiveIssue     bool         `json:"is_repetitive_issue"`      // 是否是循环事项
	IsCompleteTogether    bool         `json:"is_complete_together"`     // 是否是一并完成
	CompletesIssues       int          `json:"completes_issues"`         // 一并完成事项总数
	IssueCreateTime       int64        `json:"issue_create_time"`        // 事项创建时间
	IsImportIssueComplete bool         `json:"is_import_issue_complete"` // 导入事项时是否是完成事项
	BatchID               string       `json:"batch_id"`                 // 批量id
}

func (e CompleteIssueEvent) Name() string {
	return "complete_issue"
}

func (e CompleteIssueEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
