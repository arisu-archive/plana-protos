package protos

type BattlePassReceiveRewardResponse struct {
	ResponsePacket
	BattlePassInfo BattlePassInfoDB
	ParcelResult   ParcelResultDB
}
