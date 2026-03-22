package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentTacticResultResponse struct {
	ResponsePacket
	TacticRank                int64 `json:",omitempty,omitzero"`
	CampaignStageHistoryDB    CampaignStageHistoryDB
	LevelUpCharacterDBs       []CharacterDB
	FirstClearReward          []ParcelInfo
	StrategyObject            Strategy
	StrategyObjectRewards     *mapx.OrderedMap[int64, []ParcelInfo]
	BonusReward               []ParcelInfo
	ParcelResultDB            ParcelResultDB
	SaveDataDB                EventContentMainStageSaveDB
	EventContentCollectionDBs []EventContentCollectionDB
}
