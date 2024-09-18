package items

import "github.com/pengcainiao/zero/tools/sensors"

type Workspace struct {
	ItemID                           string `json:"item_id"`                                        // 空间ID
	ItemType                         string `json:"item_type"`                                      // 类型
	NumberOfSpaceParticipants        int    `json:"number_of_space_participants,omitempty"`         // 空间成员数
	NumberOfSpaceOutsideParticipants int    `json:"number_of_space_outside_participants,omitempty"` // 空间外部成员数
	NumberOfSpaceItems               int    `json:"number_of_space_items,omitempty"`                // 空间项目数
}

func NewWorkspace(itemID string) Workspace {
	return Workspace{
		ItemID:                           itemID,
		NumberOfSpaceParticipants:        -1,
		NumberOfSpaceOutsideParticipants: -1,
		NumberOfSpaceItems:               -1,
	}
}

func (Workspace) IsAnonymous() bool {
	return false
}

func (w Workspace) GetItemID() string {
	return w.ItemID
}

func (w Workspace) GetItemType() string {
	return w.Name()
}

func (w Workspace) Name() string {
	return "workspace"
}

func (w Workspace) FillLib(_ *sensors.Lib) {

}

func (w Workspace) Properties() map[string]interface{} {
	mapData := sensors.Struct2Map(w)
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
