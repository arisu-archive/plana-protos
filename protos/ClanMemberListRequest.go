package protos

type ClanMemberListRequest struct {
	RequestPacket
	ClanDBId int64 `json:",omitempty,omitzero"`
}
