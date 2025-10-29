package protos

type ClanKickRequest struct {
	RequestPacket
	MemberAccountId int64 `json:",omitempty,omitzero"`
}
