package protos

type WelcomeCampaignMissionMultipleRewardResponse struct {
	ResponsePacket
	ParcelResultDB  *ParcelResultDB `json:",omitempty,omitzero"`
	AddedHistoryDBs []*MissionHistoryDB
}
