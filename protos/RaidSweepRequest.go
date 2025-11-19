package protos

type RaidSweepRequest struct {
	RequestPacket
	UniqueId   int64 `json:",omitempty,omitzero"`
	SweepCount int64 `json:",omitempty,omitzero"`
}
