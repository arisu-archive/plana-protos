package protos

type ScenarioSelectRequest struct {
	RequestPacket
	ScriptGroupId int64 `json:",omitempty,omitzero"`
	ScriptSelectGroup int64 `json:",omitempty,omitzero"`
}
