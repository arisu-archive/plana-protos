package protos

type AccountGetTutorialRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
