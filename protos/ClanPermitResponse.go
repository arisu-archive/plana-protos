package protos

type ClanPermitResponse struct {
	ResponsePacket
	ClanDB ClanDB `json:",omitempty,omitzero"`
	ClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
}
