package protos

type ClanQuitRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
