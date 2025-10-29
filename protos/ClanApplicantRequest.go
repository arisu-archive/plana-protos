package protos

type ClanApplicantRequest struct {
	RequestPacket
	OffSet int64 `json:",omitempty,omitzero"`
}
