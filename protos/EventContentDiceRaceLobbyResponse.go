package protos

type EventContentDiceRaceLobbyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	DiceRaceDB EventContentDiceRaceDB `json:",omitempty,omitzero"`
}
