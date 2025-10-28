package protos

type EventContentEnterMainGroundStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB EventContentMainGroundStageSaveDB `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
