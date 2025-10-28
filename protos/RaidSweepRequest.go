package protos

type RaidSweepRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	SweepCount int64 `json:",omitempty,omitzero"`
}
