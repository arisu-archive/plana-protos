package protos

type TimeAttackDungeonLoginRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
