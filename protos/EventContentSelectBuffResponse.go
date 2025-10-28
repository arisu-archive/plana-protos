package protos

type EventContentSelectBuffResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
}
