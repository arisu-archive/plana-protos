package protos

type DailyRecordRestoreStreakAttendanceRequest struct {
	RequestPacket
	StreakRecordId int64 `json:",omitempty,omitzero"`
}
