package protos

type EliminateRaidRankingIndexResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RankBrackets []RaidRankBracket `json:",omitempty,omitzero"`
}
