package protos

type MemoryLobbySetMainResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountDB AccountDB `json:",omitempty,omitzero"`
}
