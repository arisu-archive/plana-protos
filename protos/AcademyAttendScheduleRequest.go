package protos

type AcademyAttendScheduleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ZoneId int64 `json:",omitempty,omitzero"`
}
