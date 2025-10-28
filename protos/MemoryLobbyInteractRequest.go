package protos

type MemoryLobbyInteractRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
