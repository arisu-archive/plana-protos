package protos

type EliminateRaidRankingRewardResponse struct {
	ResponsePacket
	ReceivedRankingRewardId int64          `json:",omitempty,omitzero"`
	ParcelResultDB          ParcelResultDB `json:",omitempty,omitzero"`
}
