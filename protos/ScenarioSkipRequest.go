package protos

type ScenarioSkipRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScriptGroupId int64 `json:",omitempty,omitzero"`
	SkipPointScriptCount int32 `json:",omitempty,omitzero"`
}
