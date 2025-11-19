package protos

type ArenaRankListRequest struct {
	RequestPacket
	StartIndex int32 `json:",omitempty,omitzero"`
	Count      int32 `json:",omitempty,omitzero"`
}
