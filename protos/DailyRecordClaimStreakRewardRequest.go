package protos

type DailyRecordClaimStreakRewardRequest struct {
	RequestPacket
	StreakRecordId int64 `json:",omitempty,omitzero"`
}
