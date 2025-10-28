package protos

type MiniGameDreamMakerAttendScheduleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	ScheduleGroupId int64 `json:",omitempty,omitzero"`
}
