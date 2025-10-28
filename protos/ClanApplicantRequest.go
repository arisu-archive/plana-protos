package protos

type ClanApplicantRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	OffSet int64 `json:",omitempty,omitzero"`
}
