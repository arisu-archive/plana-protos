package protos

type EventContentConcentrationRoundSkipResponse struct {
	ResponsePacket
	SaveDB         EventContentConcentrationSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB                  `json:",omitempty,omitzero"`
}
