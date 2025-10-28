package protos

type OpenConditionSetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConditionDBs []OpenConditionDB `json:",omitempty,omitzero"`
}
