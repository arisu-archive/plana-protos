package protos

type ClanMemberListResponse struct {
	ResponsePacket
	ClanDB        ClanDB         `json:",omitempty,omitzero"`
	ClanMemberDBs []ClanMemberDB `json:",omitempty,omitzero"`
}
