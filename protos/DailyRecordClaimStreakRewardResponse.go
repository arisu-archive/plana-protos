package protos

type DailyRecordClaimStreakRewardResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
	StreakRecordDB *StreakRecordDB `json:",omitempty,omitzero"`
}
