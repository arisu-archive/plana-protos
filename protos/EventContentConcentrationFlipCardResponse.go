package protos

type EventContentConcentrationFlipCardResponse struct {
	ResponsePacket
	SaveDB         EventContentConcentrationSaveDB `json:",omitempty,omitzero"`
	First          EventContentConcentrationCardDB `json:",omitempty,omitzero"`
	Second         EventContentConcentrationCardDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB                  `json:",omitempty,omitzero"`
}
