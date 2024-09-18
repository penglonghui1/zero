package items

import "github.com/pengcainiao/zero/tools/sensors"

const (
	ObRankSelf    = "个人级"
	ObRankDept    = "部门级"
	ObRankCompany = "公司级"

	ObStatusNormal = "正常"
	ObStatusRisk   = "风险"
	ObStatusDelay  = "延迟"
	ObStatusDone   = "完结"
)

type Objective struct {
	ItemID                       string  `json:"item_id,omitempty"`                        // 目标ID
	ItemType                     string  `json:"item_type,omitempty"`                      // 类型
	ObjectiveRank                string  `json:"objective_rank,omitempty"`                 // 目标级别
	ObjectiveStatus              string  `json:"objective_status,omitempty"`               // 目标状态
	ObjectiveProgress            float64 `json:"objective_progress,omitempty"`             // 目标进度
	KrNumber                     int     `json:"kr_number,omitempty"`                      // 目标KR数
	AlignObjectiveNumber         int     `json:"align_objective_number,omitempty"`         // 对齐目标数
	ViewNumber                   int     `json:"view_number,omitempty"`                    // 视图数
	RelevanceNumber              int     `json:"relevance_number,omitempty"`               // 关联数
	CommentNumber                int     `json:"comment_number,omitempty"`                 // 评论数
	ObjectiveParticipantQuantity int     `json:"objective_participant_quantity,omitempty"` // 目标参与人数量
}

func NewObjective(itemID string) Objective {
	return Objective{
		ItemID:                       itemID,
		ItemType:                     "objective",
		ObjectiveRank:                "-1",
		ObjectiveStatus:              "-1",
		ObjectiveProgress:            -1,
		KrNumber:                     -1,
		AlignObjectiveNumber:         -1,
		ViewNumber:                   -1,
		RelevanceNumber:              -1,
		CommentNumber:                -1,
		ObjectiveParticipantQuantity: -1,
	}
}

func (o Objective) IsAnonymous() bool {
	return false
}

func (o Objective) GetItemID() string {
	return o.ItemID
}

func (o Objective) GetItemType() string {
	return o.Name()
}

func (o Objective) Name() string {
	return "objective"
}

func (o Objective) FillLib(_ *sensors.Lib) {

}

func (o Objective) Properties() map[string]interface{} {
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
