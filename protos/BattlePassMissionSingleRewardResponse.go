package protos

type BattlePassMissionSingleRewardResponse struct {
	ResponsePacket
	AddedHistoryDB MissionHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
}
