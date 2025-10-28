package protos

type EliminateRaidLoginRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
