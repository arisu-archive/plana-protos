package protos

type EventContentTacticResultResponse struct {
	ResponsePacket
	TacticRank int64 `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs []CharacterDB `json:",omitempty,omitzero"`
	FirstClearReward []ParcelInfo `json:",omitempty,omitzero"`
	StrategyObject Strategy `json:",omitempty,omitzero"`
	StrategyObjectRewards map[int64][]ParcelInfo `json:",omitempty,omitzero"`
	BonusReward []ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
