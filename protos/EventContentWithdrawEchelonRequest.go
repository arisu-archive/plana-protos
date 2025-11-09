package protos

type EventContentWithdrawEchelonRequest struct {
	RequestPacket
	EventContentId          int64   `json:",omitempty,omitzero"`
	StageUniqueId           int64   `json:",omitempty,omitzero"`
	WithdrawEchelonEntityId []int64 `json:",omitempty,omitzero"`
}
