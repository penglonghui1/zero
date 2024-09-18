package items

import (
	"github.com/pengcainiao/zero/tools/sensors"
)

type MeetingState string

const (
	MeetingStateInit     = MeetingState("未开始")
	MeetingStateDoing    = MeetingState("进行中")
	MeetingStateFinished = MeetingState("已完成")
)

type Meeting struct {
	ItemID                         string       `json:"item_id"`
	ItemType                       string       `json:"item_type"`                                     // 这个值不用不需要设置
	MeetingParticipantQuantity     int          `json:"meeting_participant_quantity,omitempty"`        // 会议参与人数量
	MeetingRecipientQuantity       int          `json:"meeting_recipient_quantity,omitempty"`          // 会议接受人数量
	MeetingFileQuantity            int          `json:"meeting_file_quantity,omitempty"`               // 会议内文件数量
	MeetingFileSize                float32      `json:"meeting_file_size,omitempty"`                   // 会议内文件大小
	MeetingState                   MeetingState `json:"meeting_state,omitempty"`                       // 未开始、进行中、已完成
	MeetingSummaryNumber           int          `json:"meeting_summary_number,omitempty"`              // 会议结论数量
	MeetingRecordNumber            int          `json:"meeting_record_number,omitempty"`               // 会议记录数量
	AddMeetingRecordNumberOfPeople int          `json:"add_meeting_record_number_of_people,omitempty"` // 添加会议记录人数
}

func NewMeeting(itemID string) Meeting {
	return Meeting{
		ItemID:                         itemID,
		MeetingParticipantQuantity:     -1,
		MeetingRecipientQuantity:       -1,
		MeetingFileQuantity:            -1,
		MeetingFileSize:                -1,
		MeetingState:                   MeetingStateInit,
		MeetingSummaryNumber:           -1,
		MeetingRecordNumber:            -1,
		AddMeetingRecordNumberOfPeople: -1,
	}
}

func (meeting Meeting) IsAnonymous() bool {
	return false
}

func (meeting Meeting) GetItemID() string {
	return meeting.ItemID
}

func (meeting Meeting) GetItemType() string {
	return meeting.Name()
}

func (meeting Meeting) Name() string {
	return "meeting"
}

func (meeting Meeting) FillLib(_ *sensors.Lib) {
}

func (meeting Meeting) Properties() map[string]interface{} {
	mapData := sensors.Struct2Map(meeting)
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
