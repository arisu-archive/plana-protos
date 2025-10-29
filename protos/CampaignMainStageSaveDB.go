package protos

type CampaignMainStageSaveDB struct {
	ContentSaveDB
	CampaignState CampaignState `json:",omitempty,omitzero"`
	CurrentTurn int32 `json:",omitempty,omitzero"`
	EnemyClearCount int32 `json:",omitempty,omitzero"`
	LastEnemyEntityId int32 `json:",omitempty,omitzero"`
	TacticRankSCount int32 `json:",omitempty,omitzero"`
	EnemyInfos map[int64]HexaUnit `json:",omitempty,omitzero"`
	EchelonInfos map[int64]HexaUnit `json:",omitempty,omitzero"`
	WithdrawInfos map[int64][]int64 `json:",omitempty,omitzero"`
	StrategyObjects map[int64]Strategy `json:",omitempty,omitzero"`
	StrategyObjectRewards map[int64][]ParcelInfo `json:",omitempty,omitzero"`
	StrategyObjectHistory []int64 `json:",omitempty,omitzero"`
	ActivatedHexaEventsAndConditions map[int64][]int64 `json:",omitempty,omitzero"`
	HexaEventDelayedExecutions map[int64][]int64 `json:",omitempty,omitzero"`
	TileMapStates map[int32]HexaTileState `json:",omitempty,omitzero"`
	DisplayInfos []HexaDisplayInfo `json:",omitempty,omitzero"`
	DeployedEchelonInfos []HexaUnit `json:",omitempty,omitzero"`
}
