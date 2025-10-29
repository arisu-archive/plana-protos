package protos

type MemoryLobbySetMainResponse struct {
	ResponsePacket
	AccountDB AccountDB `json:",omitempty,omitzero"`
}
