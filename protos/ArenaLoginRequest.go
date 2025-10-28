package protos

type ArenaLoginRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
