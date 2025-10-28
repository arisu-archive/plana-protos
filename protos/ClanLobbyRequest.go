package protos

type ClanLobbyRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
