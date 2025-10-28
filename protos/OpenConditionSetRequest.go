package protos

type OpenConditionSetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConditionDB OpenConditionDB `json:",omitempty,omitzero"`
}
