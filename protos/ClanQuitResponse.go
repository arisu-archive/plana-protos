package protos

type ClanQuitResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
