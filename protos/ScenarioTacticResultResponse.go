package protos

type ScenarioTacticResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StrategyObject Strategy `json:",omitempty,omitzero"`
	SaveDataDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	IsPlayerWin bool `json:",omitempty,omitzero"`
	ScenarioIds []int64 `json:",omitempty,omitzero"`
}
