package protos

type EventContentWithdrawEchelonRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	WithdrawEchelonEntityId []int64 `json:",omitempty,omitzero"`
}
