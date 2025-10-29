package protos

type OpenConditionSetResponse struct {
	ResponsePacket
	ConditionDBs []OpenConditionDB `json:",omitempty,omitzero"`
}
