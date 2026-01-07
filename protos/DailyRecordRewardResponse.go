package protos

type DailyRecordRewardResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	DailyRecordDB  DailyRecordDB  `json:",omitempty,omitzero"`
}
