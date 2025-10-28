package protos

type EventContentEnterMainStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	IsOnSubEvent bool `json:",omitempty,omitzero"`
}
