package protos

type MemoryLobbySetMainRequest struct {
	RequestPacket
	MemoryLobbyId int64 `json:",omitempty,omitzero"`
}
