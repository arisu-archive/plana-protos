package protos

type EliminateRaidLobbyRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
