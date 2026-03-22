package protos

type ScenarioEndTurnResponse struct {
	ResponsePacket
	SaveDataDB        StoryStrategyStageSaveDB
	AccountCurrencyDB AccountCurrencyDB
	ScenarioIds       []int64
}
