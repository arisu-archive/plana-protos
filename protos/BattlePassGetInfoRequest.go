package protos

type BattlePassGetInfoRequest struct {
	RequestPacket
	BattlePassId int64 `json:",omitempty,omitzero"`
}
