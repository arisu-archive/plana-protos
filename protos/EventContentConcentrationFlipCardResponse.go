package protos

type EventContentConcentrationFlipCardResponse struct {
	ResponsePacket
	SaveDB         EventContentConcentrationSaveDB
	First          EventContentConcentrationCardDB
	Second         EventContentConcentrationCardDB
	ParcelResultDB ParcelResultDB
}
