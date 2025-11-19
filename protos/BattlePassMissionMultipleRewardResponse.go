package protos

type BattlePassMissionMultipleRewardResponse struct {
	ResponsePacket
	AddedHistoryDBs []MissionHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB  ParcelResultDB     `json:",omitempty,omitzero"`
	BattlePassInfo  BattlePassInfoDB   `json:",omitempty,omitzero"`
}
