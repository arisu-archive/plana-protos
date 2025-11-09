package protos

type StrategyClearRewardInfo struct {
	FirstClearReward        []ParcelInfo           `json:",omitempty,omitzero"`
	ThreeStarReward         []ParcelInfo           `json:",omitempty,omitzero"`
	StrategyObjectRewards   map[int64][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB          ParcelResultDB         `json:",omitempty,omitzero"`
	EventContentBonusReward []ParcelInfo           `json:",omitempty,omitzero"`
	CampaignStageHistoryDB  CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
