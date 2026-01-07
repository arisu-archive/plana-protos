package protos

type DailyRecordDB struct {
	Id                int64 `json:",omitempty,omitzero"`
	AttendanceDay     int32 `json:",omitempty,omitzero"`
	ReceivedRewardDay int32 `json:",omitempty,omitzero"`
}
