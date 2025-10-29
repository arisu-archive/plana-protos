package protos

type TimeAttackDungeonLoginResponse struct {
	ResponsePacket
	PreviousRoomDB TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
}
