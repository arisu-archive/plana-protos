package protos

type BattlePassMissionSingleRewardRequest struct {
	RequestPacket
	BattlePassId int64 `json:",omitempty,omitzero"`
	MissionUniqueId int64 `json:",omitempty,omitzero"`
}
