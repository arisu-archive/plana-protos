package protos

type TimeAttackDungeonLobbyRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
