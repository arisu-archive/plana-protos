package protos

type EventContentEnterStoryStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     EventContentStoryStageSaveDB
}
