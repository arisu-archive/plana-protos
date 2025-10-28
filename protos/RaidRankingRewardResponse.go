package protos

type RaidRankingRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ReceivedRankingRewardId int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
