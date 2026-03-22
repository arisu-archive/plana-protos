package protos

type WeekDungeonEnterBattleResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	Seed           int32 `json:",omitempty,omitzero"`
	Sequence       int32 `json:",omitempty,omitzero"`
}
