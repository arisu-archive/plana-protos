package protos

type MemoryLobbyUpdateLobbyModeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IsMemoryLobbyMode bool `json:",omitempty,omitzero"`
}
