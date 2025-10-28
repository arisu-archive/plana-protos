package protos

type CafeTravelRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetAccountId *int64 `json:",omitempty,omitzero"`
	CurrentVisitingAccountId *int64 `json:",omitempty,omitzero"`
}
