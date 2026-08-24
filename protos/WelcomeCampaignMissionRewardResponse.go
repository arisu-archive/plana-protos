package protos

type WelcomeCampaignMissionRewardResponse struct {
	ResponsePacket
	AddedHistoryDB *MissionHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB *ParcelResultDB   `json:",omitempty,omitzero"`
}
