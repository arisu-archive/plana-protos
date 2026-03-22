package protos

type CampaignEnterSubStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     CampaignSubStageSaveDB
}
