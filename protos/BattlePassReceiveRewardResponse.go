package protos

type BattlePassReceiveRewardResponse struct {
	ResponsePacket
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
	ParcelResult   ParcelResultDB   `json:",omitempty,omitzero"`
}
