package protos

type ArenaEnterLobbyRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
