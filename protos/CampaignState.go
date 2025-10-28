package protos

type CampaignState int32

const (
	CampaignState_BeforeStart CampaignState = 0
	CampaignState_BeginPlayerPhase CampaignState = 1
	CampaignState_PlayerPhase CampaignState = 2
	CampaignState_EndPlayerPhase CampaignState = 3
	CampaignState_BeginEnemyPhase CampaignState = 4
	CampaignState_EnemyPhase CampaignState = 5
	CampaignState_EndEnemyPhase CampaignState = 6
	CampaignState_Win CampaignState = 7
	CampaignState_Lose CampaignState = 8
	CampaignState_StrategySkip CampaignState = 9
)
