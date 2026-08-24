package protos

type DailyRecordSupplementSeasonAttendanceRequest struct {
	RequestPacket
	SeasonRecordId int64 `json:",omitempty,omitzero"`
	SupplementDays int32 `json:",omitempty,omitzero"`
}
