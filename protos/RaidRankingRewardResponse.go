package protos

type RaidRankingRewardResponse struct {
	ResponsePacket
	ReceivedRankingRewardId int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
