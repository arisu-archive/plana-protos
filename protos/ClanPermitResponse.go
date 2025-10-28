package protos

type ClanPermitResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanDB ClanDB `json:",omitempty,omitzero"`
	ClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
}
