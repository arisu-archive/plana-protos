package protos

type CampaignRestartMainStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     CampaignMainStageSaveDB
}
