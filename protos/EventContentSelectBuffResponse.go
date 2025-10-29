package protos

type EventContentSelectBuffResponse struct {
	ResponsePacket
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
}
