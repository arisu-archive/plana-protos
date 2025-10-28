package protos

type EventContentEnterTacticResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
