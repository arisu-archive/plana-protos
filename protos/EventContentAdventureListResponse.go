package protos

type EventContentAdventureListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageHistoryDBs []CampaignStageHistoryDB `json:",omitempty,omitzero"`
	StrategyObjecthistoryDBs []StrategyObjectHistoryDB `json:",omitempty,omitzero"`
	EventContentBonusRewardDBs []EventContentBonusRewardDB `json:",omitempty,omitzero"`
	AlreadyReceiveRewardId []int64 `json:",omitempty,omitzero"`
	StagePoint int64 `json:",omitempty,omitzero"`
}
