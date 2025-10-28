package protos

type ClanAutoJoinRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
