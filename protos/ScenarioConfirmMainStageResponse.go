package protos

type ScenarioConfirmMainStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     StoryStrategyStageSaveDB
	ScenarioIds    []int64
}
