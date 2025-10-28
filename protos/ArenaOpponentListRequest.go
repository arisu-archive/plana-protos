package protos

type ArenaOpponentListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
