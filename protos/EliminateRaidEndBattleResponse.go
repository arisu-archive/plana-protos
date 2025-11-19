package protos

type EliminateRaidEndBattleResponse struct {
	ResponsePacket
	RankingPoint        int64          `json:",omitempty,omitzero"`
	BestRankingPoint    int64          `json:",omitempty,omitzero"`
	ClearTimePoint      int64          `json:",omitempty,omitzero"`
	HPPercentScorePoint int64          `json:",omitempty,omitzero"`
	DefaultClearPoint   int64          `json:",omitempty,omitzero"`
	ParcelResultDB      ParcelResultDB `json:",omitempty,omitzero"`
}
