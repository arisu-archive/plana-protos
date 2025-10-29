package protos

type SchoolDungeonEnterBattleRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
