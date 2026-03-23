package protos

type ScenarioPortalResponse struct {
	ResponsePacket
	StoryStrategyStageSaveDB *StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	ScenarioIds              []int64
}
