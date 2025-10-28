package protos

type ClanMemberRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanDBId int64 `json:",omitempty,omitzero"`
	MemberAccountId int64 `json:",omitempty,omitzero"`
}
