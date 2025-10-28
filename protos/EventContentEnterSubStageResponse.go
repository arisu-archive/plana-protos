package protos

type EventContentEnterSubStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB EventContentSubStageSaveDB `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
