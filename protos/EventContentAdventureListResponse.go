package protos

type EventContentAdventureListResponse struct {
	ResponsePacket
	StageHistoryDBs            []*CampaignStageHistoryDB
	StrategyObjecthistoryDBs   []*StrategyObjectHistoryDB
	EventContentBonusRewardDBs []*EventContentBonusRewardDB
	AlreadyReceiveRewardId     []int64
	StagePoint                 int64 `json:",omitempty,omitzero"`
}
