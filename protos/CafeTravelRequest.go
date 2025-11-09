package protos

type CafeTravelRequest struct {
	RequestPacket
	TargetAccountId          *int64 `json:",omitempty,omitzero"`
	CurrentVisitingAccountId *int64 `json:",omitempty,omitzero"`
}
