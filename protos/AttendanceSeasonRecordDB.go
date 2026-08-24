package protos

type AttendanceSeasonRecordDB struct {
	SeasonRecordId    int64 `json:",omitempty,omitzero"`
	CurrentDay        int32 `json:",omitempty,omitzero"`
	ReceivedRewardDay int32 `json:",omitempty,omitzero"`
}
