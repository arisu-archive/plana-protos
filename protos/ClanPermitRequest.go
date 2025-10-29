package protos

type ClanPermitRequest struct {
	RequestPacket
	ApplicantAccountId int64 `json:",omitempty,omitzero"`
	IsPerMit bool `json:",omitempty,omitzero"`
}
