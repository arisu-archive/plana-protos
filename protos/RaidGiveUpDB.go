package protos

type RaidGiveUpDB struct {
	Ranking          int64 `json:",omitempty,omitzero"`
	RankingPoint     int64 `json:",omitempty,omitzero"`
	BestRankingPoint int64 `json:",omitempty,omitzero"`
}
