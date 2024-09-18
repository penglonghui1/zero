package events

import "github.com/pengcainiao2/zero/tools/sensors"

// BusinessType 事项类型
type BusinessType string

// UserIdentity 参与人身份
type UserIdentity string

const (
	Matter         BusinessType = "事项"
	Meeting        BusinessType = "会议"
	TimeCollect    BusinessType = "时间征集"
	Notice         BusinessType = "公告"
	Todo           BusinessType = "待办"
	Follower       UserIdentity = "协作人"
	Organizer      UserIdentity = "发起人"
	ProjectCreator UserIdentity = "项目创建人"
)

type IssueNewFileEvent struct {
	sensors.BaseEvent
	SpaceID         string       `json:"space_id"`          // 空间id
	ProjectID       string       `json:"project_id"`        // 项目id
	BusinessID      string       `json:"business_id"`       // 事项id
	BusinessType    BusinessType `json:"business_type"`     // 事项类型
	FileId          string       `json:"file_id"`           // 文件id
	FileName        string       `json:"file_name"`         // 文件名称
	FileCreatorID   string       `json:"file_creator_id"`   // 创建人id
	FileCreatorType UserIdentity `json:"file_creator_type"` // 创建人类型
	FileExtension   string       `json:"file_extension"`    // 文件后缀名
	FileSize        float32      `json:"file_size"`         // 文件大小
	FileUpload      string       `json:"file_upload"`       // 文件上传
	FileOrigin      string       `json:"file_origin"`       // 文件来源
	EnterFilePage   string       `json:"enter_file_page"`   // 进入文件页面，好像废弃？
}

func (e IssueNewFileEvent) Name() string {
	return "new_file"
}

func (e IssueNewFileEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
