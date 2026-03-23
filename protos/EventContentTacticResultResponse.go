package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentTacticResultResponse struct {
	ResponsePacket
	TacticRank                int64                   `json:",omitempty,omitzero"`
	CampaignStageHistoryDB    *CampaignStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs       []*CharacterDB
	FirstClearReward          []*ParcelInfo
	StrategyObject            *Strategy `json:",omitempty,omitzero"`
	StrategyObjectRewards     *mapx.OrderedMap[int64, []*ParcelInfo]
	BonusReward               []*ParcelInfo
	ParcelResultDB            *ParcelResultDB              `json:",omitempty,omitzero"`
	SaveDataDB                *EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []*EventContentCollectionDB
}
