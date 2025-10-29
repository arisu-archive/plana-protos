package protos

type EventContentEnterMainStageResponse struct {
	ResponsePacket
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	IsOnSubEvent bool `json:",omitempty,omitzero"`
}
