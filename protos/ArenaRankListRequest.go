package protos

type ArenaRankListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StartIndex int32 `json:",omitempty,omitzero"`
	Count int32 `json:",omitempty,omitzero"`
}
