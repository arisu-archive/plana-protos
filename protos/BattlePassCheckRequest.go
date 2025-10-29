package protos

type BattlePassCheckRequest struct {
	RequestPacket
	BattlePassId int64 `json:",omitempty,omitzero"`
}
