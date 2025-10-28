package protos

type EventContentConcentrationGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDB EventContentConcentrationSaveDB `json:",omitempty,omitzero"`
}
