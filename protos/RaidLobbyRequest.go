package protos

type RaidLobbyRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
