package protos

import (
	"github.com/arisu-archive/mapx"
)

type StrategyClearRewardInfo struct {
	FirstClearReward        []ParcelInfo
	ThreeStarReward         []ParcelInfo
	StrategyObjectRewards   *mapx.OrderedMap[int64, []ParcelInfo]
	ParcelResultDB          ParcelResultDB
	EventContentBonusReward []ParcelInfo
	CampaignStageHistoryDB  CampaignStageHistoryDB
}
