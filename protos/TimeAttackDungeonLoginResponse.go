package protos

type TimeAttackDungeonLoginResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PreviousRoomDB TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
}
