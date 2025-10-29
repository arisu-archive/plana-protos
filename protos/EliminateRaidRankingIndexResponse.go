package protos

type EliminateRaidRankingIndexResponse struct {
	ResponsePacket
	RankBrackets []RaidRankBracket `json:",omitempty,omitzero"`
}
