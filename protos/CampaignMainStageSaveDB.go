package protos

import (
	"github.com/arisu-archive/mapx"
)

type CampaignMainStageSaveDB struct {
	ContentSaveDB
	CampaignState                    CampaignState                          `json:",omitempty,omitzero"`
	CurrentTurn                      int32                                  `json:",omitempty,omitzero"`
	EnemyClearCount                  int32                                  `json:",omitempty,omitzero"`
	LastEnemyEntityId                int32                                  `json:",omitempty,omitzero"`
	TacticRankSCount                 int32                                  `json:",omitempty,omitzero"`
	EnemyInfos                       *mapx.OrderedMap[int64, HexaUnit]      `json:",omitempty,omitzero"`
	EchelonInfos                     *mapx.OrderedMap[int64, HexaUnit]      `json:",omitempty,omitzero"`
	WithdrawInfos                    *mapx.OrderedMap[int64, []int64]       `json:",omitempty,omitzero"`
	StrategyObjects                  *mapx.OrderedMap[int64, Strategy]      `json:",omitempty,omitzero"`
	StrategyObjectRewards            *mapx.OrderedMap[int64, []ParcelInfo]  `json:",omitempty,omitzero"`
	StrategyObjectHistory            []int64                                `json:",omitempty,omitzero"`
	ActivatedHexaEventsAndConditions *mapx.OrderedMap[int64, []int64]       `json:",omitempty,omitzero"`
	HexaEventDelayedExecutions       *mapx.OrderedMap[int64, []int64]       `json:",omitempty,omitzero"`
	TileMapStates                    *mapx.OrderedMap[int32, HexaTileState] `json:",omitempty,omitzero"`
	DisplayInfos                     []HexaDisplayInfo                      `json:",omitempty,omitzero"`
	DeployedEchelonInfos             []HexaUnit                             `json:",omitempty,omitzero"`
}
