package protos

type MiniGameCCGSweepRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	SweepCount int32 `json:",omitempty,omitzero"`
}
