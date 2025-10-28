package protos

type ScenarioRestartMainStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
}
