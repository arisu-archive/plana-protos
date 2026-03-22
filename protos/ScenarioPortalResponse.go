package protos

type ScenarioPortalResponse struct {
	ResponsePacket
	StoryStrategyStageSaveDB StoryStrategyStageSaveDB
	ScenarioIds              []int64
}
