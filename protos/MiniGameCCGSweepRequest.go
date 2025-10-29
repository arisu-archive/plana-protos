package protos

type MiniGameCCGSweepRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	SweepCount int32 `json:",omitempty,omitzero"`
}
