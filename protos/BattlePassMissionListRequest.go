package protos

type BattlePassMissionListRequest struct {
	RequestPacket
	BattlePassId int64 `json:",omitempty,omitzero"`
}
