package protos

type MiniGameCCGCompleteGameRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
