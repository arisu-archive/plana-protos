package protos

type CampaignTutorialStageResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ClearReward []ParcelInfo `json:",omitempty,omitzero"`
	FirstClearReward []ParcelInfo `json:",omitempty,omitzero"`
}
