package protos

type EventContentDiceRaceRollRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
