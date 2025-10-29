package protos

type EventContentDiceRaceLobbyResponse struct {
	ResponsePacket
	DiceRaceDB EventContentDiceRaceDB `json:",omitempty,omitzero"`
}
