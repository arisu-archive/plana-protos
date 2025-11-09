package protos

type ScenarioSkipRequest struct {
	RequestPacket
	ScriptGroupId        int64 `json:",omitempty,omitzero"`
	SkipPointScriptCount int32 `json:",omitempty,omitzero"`
}
