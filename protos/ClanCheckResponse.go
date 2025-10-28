package protos

type ClanCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
