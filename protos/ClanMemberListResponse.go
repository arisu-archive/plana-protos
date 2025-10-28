package protos

type ClanMemberListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanDB ClanDB `json:",omitempty,omitzero"`
	ClanMemberDBs []ClanMemberDB `json:",omitempty,omitzero"`
}
