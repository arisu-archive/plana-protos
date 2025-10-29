package protos

type OpenConditionSetRequest struct {
	RequestPacket
	ConditionDB OpenConditionDB `json:",omitempty,omitzero"`
}
