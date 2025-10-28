package protos

type EliminateRaidRankingIndexRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
