package protos

type EventContentConfirmMainStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	SaveDataDB     EventContentMainStageSaveDB
}
