package protos

type BattlePassReceiveRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
}
