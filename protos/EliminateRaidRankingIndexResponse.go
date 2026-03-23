package protos

type EliminateRaidRankingIndexResponse struct {
	ResponsePacket
	RankBrackets []*RaidRankBracket
}
