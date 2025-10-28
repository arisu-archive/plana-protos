package protos

type ClanApplicantResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanMemberDBs []ClanMemberDB `json:",omitempty,omitzero"`
}
