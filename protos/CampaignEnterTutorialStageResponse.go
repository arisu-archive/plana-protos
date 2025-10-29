package protos

type CampaignEnterTutorialStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB CampaignTutorialStageSaveDB `json:",omitempty,omitzero"`
}
