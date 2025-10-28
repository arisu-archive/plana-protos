package protos

type EliminateRaidSweepRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	SweepCount int32 `json:",omitempty,omitzero"`
}
