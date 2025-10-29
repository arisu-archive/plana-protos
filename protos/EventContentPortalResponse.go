package protos

type EventContentPortalResponse struct {
	ResponsePacket
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
}
