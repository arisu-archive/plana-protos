package protos

type TimeAttackDungeonSweepRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SweepCount int64 `json:",omitempty,omitzero"`
}
