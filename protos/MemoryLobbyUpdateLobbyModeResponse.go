package protos

type MemoryLobbyUpdateLobbyModeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
