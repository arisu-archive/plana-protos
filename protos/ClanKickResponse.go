package protos

type ClanKickResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
