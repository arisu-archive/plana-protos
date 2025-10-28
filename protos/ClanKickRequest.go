package protos

type ClanKickRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MemberAccountId int64 `json:",omitempty,omitzero"`
}
