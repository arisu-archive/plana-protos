package protos

type MemoryLobbyUpdateLobbyModeRequest struct {
	RequestPacket
	IsMemoryLobbyMode bool `json:",omitempty,omitzero"`
}
