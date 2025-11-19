package protos

type WeekDungeonEnterBattleRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonIndex  int64 `json:",omitempty,omitzero"`
}
