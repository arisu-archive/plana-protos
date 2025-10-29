package protos

type EventContentTreasureLobbyRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
