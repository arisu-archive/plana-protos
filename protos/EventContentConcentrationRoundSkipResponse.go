package protos

type EventContentConcentrationRoundSkipResponse struct {
	ResponsePacket
	SaveDBBefore   EventContentConcentrationSaveDB `json:",omitempty,omitzero"`
	SaveDB         EventContentConcentrationSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB                  `json:",omitempty,omitzero"`
}
