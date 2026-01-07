package protos

type DailyRecordRewardRequest struct {
	RequestPacket
	DailyRecordId int64 `json:",omitempty,omitzero"`
}
