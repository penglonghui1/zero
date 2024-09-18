package items

import "github.com/pengcainiao/zero/tools/sensors"

type Project struct {
	ItemID                     string `json:"item_id"`                      //项目ID
	ItemType                   string `json:"item_type"`                    //项目类型
	ProjectIssueNumber         int    `json:"project_issue_number"`         //项目事项总数
	ProjectParticipantQuantity int    `json:"project_participant_quantity"` //项目参与者总数
	TitleCharactersNumber      int    `json:"title_characters_number"`      //项目标题字符数
	DetailsCharactersNumber    int    `json:"details_characters_number"`    //项目详情字符数
}

func NewProject(itemID string) Project {
	return Project{
		ItemID:                     itemID,
		ProjectIssueNumber:         -1,
		ProjectParticipantQuantity: -1,
		TitleCharactersNumber:      -1,
		DetailsCharactersNumber:    -1,
	}
}

func (Project) IsAnonymous() bool {
	return false
}

func (p Project) GetItemID() string {
	return p.ItemID
}

func (p Project) GetItemType() string {
	return p.Name()
}

func (p Project) Name() string {
	return "project"
}

func (p Project) FillLib(_ *sensors.Lib) {

}

func (p Project) Properties() map[string]interface{} {
	mapData := sensors.Struct2Map(p)
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
