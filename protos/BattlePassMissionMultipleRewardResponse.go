package protos

type BattlePassMissionMultipleRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AddedHistoryDBs []MissionHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	BattlePassInfo BattlePassInfoDB `json:",omitempty,omitzero"`
}
