package protos

type CampaignDeployEchelonResponse struct {
	ResponsePacket
	SaveDataDB CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
