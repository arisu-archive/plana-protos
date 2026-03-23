package protos

import (
	"github.com/arisu-archive/mapx"
)

type CampaignTacticResultResponse struct {
	ResponsePacket
	TacticRank             int64                   `json:",omitempty,omitzero"`
	CampaignStageHistoryDB *CampaignStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs    []CharacterDB
	FirstClearReward       []ParcelInfo
	ThreeStarReward        []ParcelInfo
	StrategyObject         *Strategy `json:",omitempty,omitzero"`
	StrategyObjectRewards  *mapx.OrderedMap[int64, []ParcelInfo]
	ParcelResultDB         *ParcelResultDB          `json:",omitempty,omitzero"`
	SaveDataDB             *CampaignMainStageSaveDB `json:",omitempty,omitzero"`
}
