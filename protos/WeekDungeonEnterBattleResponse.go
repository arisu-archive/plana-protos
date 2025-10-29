package protos

type WeekDungeonEnterBattleResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	Seed int32 `json:",omitempty,omitzero"`
	Sequence int32 `json:",omitempty,omitzero"`
}
