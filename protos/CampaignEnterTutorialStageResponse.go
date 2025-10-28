package protos

type CampaignEnterTutorialStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB CampaignTutorialStageSaveDB `json:",omitempty,omitzero"`
}
