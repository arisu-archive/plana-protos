package protos

type RaidRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RankingPoint int64 `json:",omitempty,omitzero"`
	BestRankingPoint int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
