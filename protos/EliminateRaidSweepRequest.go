package protos

type EliminateRaidSweepRequest struct {
	RequestPacket
	UniqueId int64 `json:",omitempty,omitzero"`
	SweepCount int32 `json:",omitempty,omitzero"`
}
