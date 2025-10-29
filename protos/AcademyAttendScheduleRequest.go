package protos

type AcademyAttendScheduleRequest struct {
	RequestPacket
	ZoneId int64 `json:",omitempty,omitzero"`
}
