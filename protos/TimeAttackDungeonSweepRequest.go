package protos

type TimeAttackDungeonSweepRequest struct {
	RequestPacket
	SweepCount int64 `json:",omitempty,omitzero"`
}
