package protos

type AttendanceRewardRequest struct {
	RequestPacket
	DayByBookUniqueId      map[int64]int64 `json:",omitempty,omitzero"`
	AttendanceBookUniqueId int64           `json:",omitempty,omitzero"`
	Day                    int64           `json:",omitempty,omitzero"`
}
