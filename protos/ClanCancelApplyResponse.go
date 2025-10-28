package protos

type ClanCancelApplyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
