package protos

type EventContentDeployEchelonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
}
