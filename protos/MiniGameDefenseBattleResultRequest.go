package protos

type MiniGameDefenseBattleResultRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageId int64 `json:",omitempty,omitzero"`
	Multiplier int32 `json:",omitempty,omitzero"`
	IsPlayerWin bool `json:",omitempty,omitzero"`
	BaseDamage int32 `json:",omitempty,omitzero"`
	HeroCount int32 `json:",omitempty,omitzero"`
	AliveCount int32 `json:",omitempty,omitzero"`
	ClearSecond int32 `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
}
