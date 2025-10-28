package protos

type RaidRankingIndexResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RankBrackets []RaidRankBracket `json:",omitempty,omitzero"`
}
