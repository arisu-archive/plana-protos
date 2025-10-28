package protos

type ClanCheckRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
