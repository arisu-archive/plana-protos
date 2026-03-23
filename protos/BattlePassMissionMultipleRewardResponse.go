package protos

type BattlePassMissionMultipleRewardResponse struct {
	ResponsePacket
	AddedHistoryDBs []*MissionHistoryDB
	ParcelResultDB  *ParcelResultDB   `json:",omitempty,omitzero"`
	BattlePassInfo  *BattlePassInfoDB `json:",omitempty,omitzero"`
}
