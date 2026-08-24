package protos

type DailyRecordRestoreStreakAttendanceResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
	StreakRecordDB *StreakRecordDB `json:",omitempty,omitzero"`
}
