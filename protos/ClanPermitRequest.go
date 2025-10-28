package protos

type ClanPermitRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ApplicantAccountId int64 `json:",omitempty,omitzero"`
	IsPerMit bool `json:",omitempty,omitzero"`
}
