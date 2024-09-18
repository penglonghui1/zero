package events

import "github.com/pengcainiao/zero/tools/sensors"

type IssueRemoveFileEvent struct {
	sensors.BaseEvent
	BusinessId   string       `json:"business_id"`
	BusinessType BusinessType `json:"business_type"`
	FileId       string       `json:"file_id"`
	CommentId    string       `json:"comment_id"`
}

func (e IssueRemoveFileEvent) Name() string {
	return "remove_file"
}

func (e IssueRemoveFileEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
