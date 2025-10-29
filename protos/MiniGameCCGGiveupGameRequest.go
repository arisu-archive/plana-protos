package protos

type MiniGameCCGGiveupGameRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
