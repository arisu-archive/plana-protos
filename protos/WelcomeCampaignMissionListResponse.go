package protos

type WelcomeCampaignMissionListResponse struct {
	ResponsePacket
	MissionHistoryUniqueIds []int64
	ProgressDBs             []*MissionProgressDB
}
