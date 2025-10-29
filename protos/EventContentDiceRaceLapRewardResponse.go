package protos

type EventContentDiceRaceLapRewardResponse struct {
	ResponsePacket
	DiceRaceDB EventContentDiceRaceDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
