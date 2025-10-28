package protos

type ClanMemberListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanDBId int64 `json:",omitempty,omitzero"`
}
