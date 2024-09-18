package items

import "github.com/pengcainiao2/zero/tools/sensors"

type Application struct {
	ItemID                         string `json:"item_id,omitempty"`                          // 应用ID
	ItemType                       string `json:"item_type,omitempty"`                        // 类型
	FlowStepCountTotal             int    `json:"flow_step_count_total,omitempty"`            // 工作流步骤数
	ApplicationParticipantQuantity int    `json:"application_participant_quantity,omitempty"` // 应用参与人数量
}

func NewApplication(itemID string) Application {
	return Application{
		ItemID:                         itemID,
		ItemType:                       "application",
		FlowStepCountTotal:             -1,
		ApplicationParticipantQuantity: -1,
	}
}

func (o Application) IsAnonymous() bool {
	return false
}

func (o Application) GetItemID() string {
	return o.ItemID
}

func (o Application) GetItemType() string {
	return o.Name()
}

func (o Application) Name() string {
	return "application"
}

func (o Application) FillLib(_ *sensors.Lib) {

}

func (o Application) Properties() map[string]interface{} {
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
