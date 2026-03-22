package protos

type ScenarioRestartMainStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     StoryStrategyStageSaveDB
}
