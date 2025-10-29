package protos

type BattlePassBuyLevelRequest struct {
	RequestPacket
	BattlePassId int64 `json:",omitempty,omitzero"`
	BattlePassBuyLevelCount int32 `json:",omitempty,omitzero"`
}
