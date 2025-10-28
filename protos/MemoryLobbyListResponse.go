package protos

type MemoryLobbyListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MemoryLobbyDBs []MemoryLobbyDB `json:",omitempty,omitzero"`
}
