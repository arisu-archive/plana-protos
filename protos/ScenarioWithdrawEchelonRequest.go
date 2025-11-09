package protos

type ScenarioWithdrawEchelonRequest struct {
	RequestPacket
	StageUniqueId           int64   `json:",omitempty,omitzero"`
	WithdrawEchelonEntityId []int64 `json:",omitempty,omitzero"`
}
