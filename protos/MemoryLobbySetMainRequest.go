package protos

type MemoryLobbySetMainRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MemoryLobbyId int64 `json:",omitempty,omitzero"`
}
