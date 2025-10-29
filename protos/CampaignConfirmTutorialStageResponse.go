package protos

type CampaignConfirmTutorialStageResponse struct {
	ResponsePacket
	SaveDataDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
