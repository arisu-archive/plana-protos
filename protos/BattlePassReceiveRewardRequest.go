package protos

type BattlePassReceiveRewardRequest struct {
	RequestPacket
	BattlePassId int64 `json:",omitempty,omitzero"`
}
