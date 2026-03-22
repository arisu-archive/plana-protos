package protos

type EventContentConcentrationRoundCompleteResponse struct {
	ResponsePacket
	SaveDBBefore   EventContentConcentrationSaveDB
	SaveDB         EventContentConcentrationSaveDB
	ParcelResultDB ParcelResultDB
}
