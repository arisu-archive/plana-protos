package protos

type MissionMultipleRewardResponse struct {
	ResponsePacket
	AddedHistoryDBs []*MissionHistoryDB
	ParcelResultDB  *ParcelResultDB `json:",omitempty,omitzero"`
}
