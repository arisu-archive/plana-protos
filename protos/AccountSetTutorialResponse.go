package protos

type AccountSetTutorialResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
