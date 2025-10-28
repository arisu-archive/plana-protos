package protos

type MemoryLobbyInteractResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
