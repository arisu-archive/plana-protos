package protos

import (
	"github.com/arisu-archive/mapx"
)

type StrategyClearRewardInfo struct {
	FirstClearReward        []*ParcelInfo
	ThreeStarReward         []*ParcelInfo
	StrategyObjectRewards   *mapx.OrderedMap[int64, []*ParcelInfo]
	ParcelResultDB          *ParcelResultDB `json:",omitempty,omitzero"`
	EventContentBonusReward []*ParcelInfo
	CampaignStageHistoryDB  *CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
