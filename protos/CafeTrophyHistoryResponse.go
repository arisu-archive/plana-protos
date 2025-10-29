package protos

type CafeTrophyHistoryResponse struct {
	ResponsePacket
	RaidSeasonRankingHistoryDBs []RaidSeasonRankingHistoryDB `json:",omitempty,omitzero"`
}
