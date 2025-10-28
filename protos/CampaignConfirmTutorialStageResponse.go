package protos

type CampaignConfirmTutorialStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
