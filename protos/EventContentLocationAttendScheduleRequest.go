package protos

type EventContentLocationAttendScheduleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	ZoneId int64 `json:",omitempty,omitzero"`
	Count int64 `json:",omitempty,omitzero"`
}
