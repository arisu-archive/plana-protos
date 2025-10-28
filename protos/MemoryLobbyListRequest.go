package protos

type MemoryLobbyListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
