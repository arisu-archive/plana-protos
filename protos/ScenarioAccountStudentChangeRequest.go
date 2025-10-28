package protos

type ScenarioAccountStudentChangeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountStudent int64 `json:",omitempty,omitzero"`
	AccountStudentBefore int64 `json:",omitempty,omitzero"`
}
