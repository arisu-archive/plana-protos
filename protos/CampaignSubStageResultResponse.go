package protos

type CampaignSubStageResultResponse struct {
	ResponsePacket
	TacticRank             int64 `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB
	LevelUpCharacterDBs    []CharacterDB
	ParcelResultDB         ParcelResultDB
	FirstClearReward       []ParcelInfo
	ThreeStarReward        []ParcelInfo
}
