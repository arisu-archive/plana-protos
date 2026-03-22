package protos

type EventContentDiceRaceLapRewardResponse struct {
	ResponsePacket
	DiceRaceDB     EventContentDiceRaceDB
	ParcelResultDB ParcelResultDB
}
