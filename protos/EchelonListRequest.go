package protos

type EchelonListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
