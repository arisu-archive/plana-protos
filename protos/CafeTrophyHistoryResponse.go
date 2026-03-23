package protos

type CafeTrophyHistoryResponse struct {
	ResponsePacket
	RaidSeasonRankingHistoryDBs []*RaidSeasonRankingHistoryDB
}
