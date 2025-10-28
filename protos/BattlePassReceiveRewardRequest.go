package protos

type BattlePassReceiveRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BattlePassId int64 `json:",omitempty,omitzero"`
}
