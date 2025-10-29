package protos

type CampaignEnterMainStageResponse struct {
	ResponsePacket
	SaveDataDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
