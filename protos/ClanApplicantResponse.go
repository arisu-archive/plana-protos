package protos

type ClanApplicantResponse struct {
	ResponsePacket
	ClanMemberDBs []ClanMemberDB `json:",omitempty,omitzero"`
}
