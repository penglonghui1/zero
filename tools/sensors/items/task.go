package items

import "github.com/pengcainiao/zero/tools/sensors"

const (
	HierarchyTypeParent = "父级"
	HierarchyTypeSon    = "子级"
	HierarchyTypeSun    = "孙级"
)

type Task struct {
	ItemID                         string  `json:"item_id"`
	ItemType                       string  `json:"item_type"`                                    // 这个值不用不需要设置
	ItemParticipantQuantity        int     `json:"item_participant_quantity,omitempty"`          // 事项参与人数量
	ItemRecipientQuantity          int     `json:"item_recipient_quantity,omitempty"`            // 事项接受人数量
	ItemFileQuantity               int     `json:"item_file_quantity,omitempty"`                 // 事项内文件数量
	ItemFileSize                   float32 `json:"item_file_size,omitempty"`                     // 事项内文件大小
	UndoneNumberOfPeople           int     `json:"undone_number_of_people,omitempty"`            // 未完成人数
	CompletedNumberOfPeople        int     `json:"completed_number_of_people,omitempty"`         // 已完成人数
	PostponedNumberOfPeople        int     `json:"postponed_number_of_people,omitempty"`         // 已延期人数
	PostponeCompleteNumberOfPeople int     `json:"postpone_complete_number_of_people,omitempty"` // 延期完成人数
	ItemEmphasisNumber             int     `json:"item_emphasis_number,omitempty"`               // 事项重点数量
	ItemRecordNumber               int     `json:"item_record_number,omitempty"`                 // 事项记录数量
	AddItemRecordNumberOfPeople    int     `json:"add_item_record_number_of_people,omitempty"`   // 添加事项记录人数
	IsAddLabel                     bool    `json:"is_add_label,omitempty"`                       // 是否添加标签
	LabelNumber                    int     `json:"label_number,omitempty"`                       // 标签数量
	IssueHierarchyType             string  `json:"issue_hierarchy_type,omitempty"`               // 事项层级类型(父级、子级、孙级)
	IssueTitle                     string  `json:"issue_title,omitempty"`                        // 事项标题
	IsStartAlarmTime               bool    `json:"is_start_alarm_time,omitempty"`                // 是否有开始时间
	HasDeadline                    bool    `json:"has_deadline,omitempty"`                       // 是否有截止时间
	IsAlarmDeadline                bool    `json:"is_alarm_deadline,omitempty"`                  // 是否有提醒时间
	IsRepetitiveIssue              bool    `json:"is_repetitive_issue,omitempty"`                // 是否是循环事项
	RefuseNumberOfPeople           int     `json:"refuse_number_of_people,omitempty"`            // 事项拒绝人数
	IsGirlCreate                   bool    `json:"is_girl_create,omitempty"`                     // 是否是小姐姐创建
	ProjectID                      string  `json:"project_id,omitempty"`                         // 项目id
	FlowStepCountTotal             int     `json:"flow_step_count_total,omitempty"`              // 工作流步骤数
}

func (task Task) IsAnonymous() bool {
	return false
}

func NewTask(itemID string) Task {
	return Task{
		ItemID:                         itemID,
		ItemParticipantQuantity:        -1,
		ItemRecipientQuantity:          -1,
		ItemFileQuantity:               -1,
		ItemFileSize:                   -1,
		UndoneNumberOfPeople:           1,
		CompletedNumberOfPeople:        -1,
		PostponedNumberOfPeople:        -1,
		PostponeCompleteNumberOfPeople: -1,
		ItemEmphasisNumber:             -1,
		ItemRecordNumber:               -1,
		AddItemRecordNumberOfPeople:    -1,
		IsAddLabel:                     false,
		LabelNumber:                    -1,
		IssueHierarchyType:             HierarchyTypeParent,
		IssueTitle:                     "-1",
		IsStartAlarmTime:               false,
		HasDeadline:                    false,
		IsAlarmDeadline:                false,
		IsRepetitiveIssue:              false,
		RefuseNumberOfPeople:           -1,
		IsGirlCreate:                   false,
		ProjectID:                      "-1",
		FlowStepCountTotal:             -1,
	}
}

func (task Task) GetItemID() string {
	return task.ItemID
}

func (task Task) GetItemType() string {
	return task.Name()
}

func (task Task) Name() string {
	return "issue"
}

func (task Task) FillLib(_ *sensors.Lib) {
}

func (task Task) Properties() map[string]interface{} {
	mapData := sensors.Struct2Map(task)
	delete(mapData, "item_id")
	delete(mapData, "item_type")

	for key, value := range mapData {
		switch val := value.(type) {
		case float64:
			if val == -1 {
				mapData[key] = 0
			}
		}
	}

	return mapData
}
