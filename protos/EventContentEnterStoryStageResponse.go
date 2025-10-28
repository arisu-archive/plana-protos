package protos

type EventContentEnterStoryStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	SaveDataDB EventContentStoryStageSaveDB `json:",omitempty,omitzero"`
}
