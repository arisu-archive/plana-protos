package protos

type RaidRankingIndexResponse struct {
	ResponsePacket
	RankBrackets []RaidRankBracket `json:",omitempty,omitzero"`
}
