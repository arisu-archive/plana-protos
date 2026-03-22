package protos

import (
	"github.com/arisu-archive/mapx"
)

type CampaignMainStageSaveDB struct {
	ContentSaveDB
	CampaignState                    CampaignState `json:",omitempty,omitzero"`
	CurrentTurn                      int32         `json:",omitempty,omitzero"`
	EnemyClearCount                  int32         `json:",omitempty,omitzero"`
	LastEnemyEntityId                int32         `json:",omitempty,omitzero"`
	TacticRankSCount                 int32         `json:",omitempty,omitzero"`
	EnemyInfos                       *mapx.OrderedMap[int64, HexaUnit]
	EchelonInfos                     *mapx.OrderedMap[int64, HexaUnit]
	WithdrawInfos                    *mapx.OrderedMap[int64, []int64]
	StrategyObjects                  *mapx.OrderedMap[int64, Strategy]
	StrategyObjectRewards            *mapx.OrderedMap[int64, []ParcelInfo]
	StrategyObjectHistory            []int64
	ActivatedHexaEventsAndConditions *mapx.OrderedMap[int64, []int64]
	HexaEventDelayedExecutions       *mapx.OrderedMap[int64, []int64]
	TileMapStates                    *mapx.OrderedMap[int32, HexaTileState]
	DisplayInfos                     []HexaDisplayInfo
	DeployedEchelonInfos             []HexaUnit
}
