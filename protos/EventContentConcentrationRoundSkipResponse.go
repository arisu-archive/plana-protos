package protos

type EventContentConcentrationRoundSkipResponse struct {
	ResponsePacket
	SaveDBBefore   EventContentConcentrationSaveDB
	SaveDB         EventContentConcentrationSaveDB
	ParcelResultDB ParcelResultDB
}
