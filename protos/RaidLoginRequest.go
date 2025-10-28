package protos

type RaidLoginRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
