package protos

type EventContentEnterStoryStageResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB EventContentStoryStageSaveDB `json:",omitempty,omitzero"`
}
