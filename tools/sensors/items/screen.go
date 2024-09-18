package items

import "github.com/pengcainiao2/zero/tools/sensors"

type Screen struct {
	ItemID          string `json:"item_id,omitempty"`          // 视图ID
	ItemType        string `json:"item_type,omitempty"`        // 类型
	IssueNumber     int    `json:"issue_number,omitempty"`     // 事项数
	ProjectNumber   int    `json:"project_number,omitempty"`   // 项目数
	ObjectiveNumber int    `json:"objective_number,omitempty"` // 目标数
	NoteNumber      int    `json:"note_number,omitempty"`      // 笔记数
	FileNumber      int    `json:"file_number,omitempty"`      // 文件数
	TextNumber      int    `json:"text_number,omitempty"`      // 文本数
}

func NewScreen(itemID string) Screen {
	return Screen{
		ItemID:          itemID,
		ItemType:        "screen",
		IssueNumber:     -1,
		ProjectNumber:   -1,
		ObjectiveNumber: -1,
		NoteNumber:      -1,
		FileNumber:      -1,
		TextNumber:      -1,
	}
}

func (o Screen) IsAnonymous() bool {
	return false
}

func (o Screen) GetItemID() string {
	return o.ItemID
}

func (o Screen) GetItemType() string {
	return o.Name()
}

func (o Screen) Name() string {
	return "screen"
}

func (o Screen) FillLib(_ *sensors.Lib) {

}

func (o Screen) Properties() map[string]interface{} {
	mapData := sensors.Struct2Map(o)
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
