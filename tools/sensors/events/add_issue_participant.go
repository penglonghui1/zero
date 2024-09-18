package events

import "github.com/pengcainiao/zero/tools/sensors"

type AddIssueParticipantEvent struct {
	sensors.BaseEvent
	IssueId                      string `json:"issue_id"`
	AddIssueParticipantType      string `json:"add_issue_participant_type"`
	InvitePeopleId               string `json:"invite_people_id"`
	EnterAddIssueParticipantPage string `json:"enter_add_issue_participant_page"`
}

func (e AddIssueParticipantEvent) Name() string {
	return "add_issue_participant"
}

func (e AddIssueParticipantEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
