package protos

type MiniGameMissionMultipleRewardResponse struct {
	ResponsePacket
	AddedHistoryDBs []MissionHistoryDB
	ParcelResultDB  *ParcelResultDB `json:",omitempty,omitzero"`
}
