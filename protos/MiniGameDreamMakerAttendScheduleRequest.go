package protos

type MiniGameDreamMakerAttendScheduleRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	ScheduleGroupId int64 `json:",omitempty,omitzero"`
}
