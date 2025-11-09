package protos

type CampaignRestartMainStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB          `json:",omitempty,omitzero"`
	SaveDataDB     CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
