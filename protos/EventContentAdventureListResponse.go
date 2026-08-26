package protos

type EventContentAdventureListResponse struct {
	ResponsePacket
	StageHistoryDBs            []*CampaignStageHistoryDB
	StrategyObjectHistoryIds   []int64
	EventContentBonusRewardDBs []*EventContentBonusRewardDB
	AlreadyReceiveRewardId     []int64
	StagePoint                 int64 `json:",omitempty,omitzero"`
}
