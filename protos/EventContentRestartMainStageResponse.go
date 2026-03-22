package protos

type EventContentRestartMainStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     EventContentMainStageSaveDB
}
