package protos

type ClanCancelApplyRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
