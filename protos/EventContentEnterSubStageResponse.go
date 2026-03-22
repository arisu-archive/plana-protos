package protos

type EventContentEnterSubStageResponse struct {
	ResponsePacket
	ParcelResultDB         ParcelResultDB
	SaveDataDB             EventContentSubStageSaveDB
	CampaignStageHistoryDB CampaignStageHistoryDB
}
