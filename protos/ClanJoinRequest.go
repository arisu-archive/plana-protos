package protos

type ClanJoinRequest struct {
	RequestPacket
	ClanDBId int64 `json:",omitempty,omitzero"`
}
