package protos

type ClanDismissResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
