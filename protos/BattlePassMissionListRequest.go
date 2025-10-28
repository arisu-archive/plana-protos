package protos

type BattlePassMissionListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BattlePassId int64 `json:",omitempty,omitzero"`
}
