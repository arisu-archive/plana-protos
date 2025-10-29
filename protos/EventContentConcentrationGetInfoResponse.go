package protos

type EventContentConcentrationGetInfoResponse struct {
	ResponsePacket
	SaveDB EventContentConcentrationSaveDB `json:",omitempty,omitzero"`
}
