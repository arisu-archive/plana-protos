package protos

type RaidRankingIndexRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
