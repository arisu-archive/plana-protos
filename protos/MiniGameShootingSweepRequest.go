package protos

type MiniGameShootingSweepRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	SweepCount int64 `json:",omitempty,omitzero"`
}
