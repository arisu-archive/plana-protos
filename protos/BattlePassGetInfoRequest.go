package protos

type BattlePassGetInfoRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BattlePassId int64 `json:",omitempty,omitzero"`
}
