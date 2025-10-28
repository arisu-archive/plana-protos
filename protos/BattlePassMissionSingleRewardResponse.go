package protos

type BattlePassMissionSingleRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AddedHistoryDB MissionHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
}
