package protos

type BattlePassGetInfoResponse struct {
	ResponsePacket
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
}
