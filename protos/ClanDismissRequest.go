package protos

type ClanDismissRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
