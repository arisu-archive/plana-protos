package protos

type CafeTrophyHistoryResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidSeasonRankingHistoryDBs []RaidSeasonRankingHistoryDB `json:",omitempty,omitzero"`
}
