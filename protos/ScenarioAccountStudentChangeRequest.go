package protos

type ScenarioAccountStudentChangeRequest struct {
	RequestPacket
	AccountStudent int64 `json:",omitempty,omitzero"`
	AccountStudentBefore int64 `json:",omitempty,omitzero"`
}
