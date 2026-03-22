package protos

import (
	"github.com/arisu-archive/mapx"
)

type CampaignTacticResultResponse struct {
	ResponsePacket
	TacticRank             int64 `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB
	LevelUpCharacterDBs    []CharacterDB
	FirstClearReward       []ParcelInfo
	ThreeStarReward        []ParcelInfo
	StrategyObject         Strategy
	StrategyObjectRewards  *mapx.OrderedMap[int64, []ParcelInfo]
	ParcelResultDB         ParcelResultDB
	SaveDataDB             CampaignMainStageSaveDB
}
