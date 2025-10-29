package protos

type MemoryLobbyListResponse struct {
	ResponsePacket
	MemoryLobbyDBs []MemoryLobbyDB `json:",omitempty,omitzero"`
}
