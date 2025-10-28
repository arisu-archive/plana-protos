package protos

type BattlePassBuyLevelRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BattlePassId int64 `json:",omitempty,omitzero"`
	BattlePassBuyLevelCount int32 `json:",omitempty,omitzero"`
}
