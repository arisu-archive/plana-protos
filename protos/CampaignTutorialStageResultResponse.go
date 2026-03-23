package protos

type CampaignTutorialStageResultResponse struct {
	ResponsePacket
	CampaignStageHistoryDB *CampaignStageHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB         *ParcelResultDB         `json:",omitempty,omitzero"`
	ClearReward            []ParcelInfo
	FirstClearReward       []ParcelInfo
}
