package protos

type ScenarioTacticResultResponse struct {
	ResponsePacket
	StrategyObject Strategy
	SaveDataDB     StoryStrategyStageSaveDB
	IsPlayerWin    bool `json:",omitempty,omitzero"`
	ScenarioIds    []int64
}
