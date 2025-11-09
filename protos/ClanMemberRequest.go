package protos

type ClanMemberRequest struct {
	RequestPacket
	ClanDBId        int64 `json:",omitempty,omitzero"`
	MemberAccountId int64 `json:",omitempty,omitzero"`
}
