package protos

type EventContentConcentrationRoundCompleteResponse struct {
	ResponsePacket
	SaveDB         EventContentConcentrationSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB                  `json:",omitempty,omitzero"`
}
