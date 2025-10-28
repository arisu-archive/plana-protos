package protos

type BattlePassGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
}
