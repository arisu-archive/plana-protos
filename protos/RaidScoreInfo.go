package protos

type RaidScoreInfo struct {
	ClearTimePoint      int64 `json:",omitempty,omitzero"`
	HPPercentScorePoint int64 `json:",omitempty,omitzero"`
	DefaultClearPoint   int64 `json:",omitempty,omitzero"`
	RankingPoint        int64 `json:",omitempty,omitzero"`
	BestRankingPoint    int64 `json:",omitempty,omitzero"`
}
